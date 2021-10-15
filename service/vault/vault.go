package vault

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/SteinsElite/pickGinS/logging"
	"github.com/ethereum/go-ethereum/common"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"

	"github.com/SteinsElite/pickGinS/internal/gateway"
	"github.com/SteinsElite/pickGinS/internal/storage"
	"github.com/SteinsElite/pickGinS/service/coin"
	"github.com/SteinsElite/pickGinS/util"
)

// package vault is used to cache the vault info from the contract(e.g. total volume,total profit)
// it will cache the data every 30 minutes and persist the data into the database at every midnight

const (
	Week  = "7D"
	Month = "1M"
	Year  = "1Y"

	collName = "vault"
)

// vault watcher will store the latest vault status since last update which
// will be queried by other parts
var cacher *CacheClient

type ValuePair struct {
	TimeStamp int64
	Value     float64
}

// Stats the vault status at the specific timestamp
type Stats struct {
	TimeStamp  int64
	CoinAmount map[string]float64
	Profit     float64 // total profit util now
}

// set up the vault watcher to provide service for querying
func initCacher() {
	cacher = &CacheClient{
		RpcClient: gateway.GetRpcClient(),
	}
	var err error
	cacher.stats, err = cacher.VaultStatsFromChain()
	if err != nil {
		log.Fatal(fmt.Errorf("fail to init cacher: %w", err))
	}
	cacher.stats.TimeStamp = time.Now().Unix()
}

type CacheClient struct {
	*gateway.RpcClient
	stats Stats
}

// get the float64 representation of the amount,if necessary, use big.float instead
func (c *CacheClient) tokenAmount(token common.Address) (fAmount float64, err error) {
	tokenState, err := c.Instance.TokenState(nil, token)
	if err != nil {
		err = fmt.Errorf("TokenState-[%w]", err)
		return
	}
	tokenVolume := new(big.Int).Sub(tokenState.Max, tokenState.Remain)
	fAmount = util.Amount2Float(tokenVolume)
	return
}

func (c *CacheClient) profitAmount() (fProfit float64, err error) {
	profit, err := c.Instance.ViewAccumulatedProfit(nil)
	if err != nil {
		err = fmt.Errorf("TokenState-[%w]", err)
		return
	}
	fProfit = util.Amount2Float(profit)
	return
}

// VaultStatsFromChain get the latest vault status from blockchain,
func (c *CacheClient) VaultStatsFromChain() (stats Stats, err error) {
	profit, err := c.profitAmount()
	btc, err := c.tokenAmount(util.BTCAddr)
	eth, err := c.tokenAmount(util.ETHAddr)
	usdt, err := c.tokenAmount(util.USDTAddr)
	ht, err := c.tokenAmount(util.HTAddr)
	mdx, err := c.tokenAmount(util.MDXAddr)
	if err != nil {
		if c.RpcClient.IsClientConnected() {
			err = fmt.Errorf("get stats from chain: %w", err)
			return
		}
		if ok, _ := c.RpcClient.SelectEndpoint(); ok {
			// change to a connected rpc client
			profit, err = c.profitAmount()
			btc, err = c.tokenAmount(util.BTCAddr)
			eth, err = c.tokenAmount(util.ETHAddr)
			usdt, err = c.tokenAmount(util.USDTAddr)
			ht, err = c.tokenAmount(util.HTAddr)
			mdx, err = c.tokenAmount(util.MDXAddr)
			if err != nil {
				err = fmt.Errorf("get stats from chain: %w", err)
				return
			}
		}
	}
	// if no error occur when interact with the blockchain, return the new status
	stats.Profit = profit
	stats.CoinAmount = map[string]float64{
		util.BTC:  btc,
		util.ETH:  eth,
		util.USDT: usdt,
		util.HT:   ht,
		util.MDX:  mdx,
	}
	return
}

// get the qualified status start since the start time
func getQualifiedStatsFromDb(startTime int64) ([]Stats, error) {
	coll := storage.AccessCollections(collName)
	opt := options.Find()
	// get the stats ascending by timestamp
	opt.SetSort(bson.D{{"timestamp", 1}})
	cur, err := coll.Find(
		context.TODO(),
		bson.D{{"timestamp", bson.D{{"$gte", startTime}}}},
		opt,
	)
	if err != nil {
		return nil, fmt.Errorf("fail find in db-[%w]", err)
	}
	var result []Stats
	err = cur.All(context.TODO(), &result)
	if err != nil {
		return nil, fmt.Errorf("fail decode data-[%w]", err)
	}
	defer cur.Close(context.TODO())
	return result, nil
}

func getQualifiedProfitFromDb(ticks []int64) ([]ValuePair, error) {
	coll := storage.AccessCollections(collName)
	var profits []ValuePair
	for _, v := range ticks {
		opt := options.Find()

		// find the latest timestamp to the time tick,
		opt.SetSort(bson.D{{"timestamp", -1}})
		opt.SetLimit(1)
		// find record that timestamp is less than or equal to the tick time
		cur, err := coll.Find(
			context.TODO(),
			bson.D{{"timestamp", bson.D{{"$lte", v}}}},
			opt,
		)
		if err != nil {
			return nil, fmt.Errorf("fail find in db-[%w]", err)
		}
		if cur.Next(context.TODO()) {
			var stats Stats
			err := cur.Decode(&stats)
			if err != nil {
				return nil, fmt.Errorf("fail decode data-[%w]", err)
			}
			profits = append(profits, ValuePair{
				TimeStamp: v,
				Value:     stats.Profit,
			})
		} else {
			// if we can't get the stats at the timestamp: we don't start the vault cacher since
			// that time, so just set it to 0
			profits = append(profits, ValuePair{
				TimeStamp: v,
				Value:     0.0,
			})
		}
	}
	return profits, nil
}

// calculate the total volume value in USD of the specific stats
func volumeValue(stats Stats) float64 {
	var totalValue float64
	for k, amount := range stats.CoinAmount {
		totalValue += amount * coin.GetCurrentCoinPrice(k)
	}
	return totalValue
}

// calculate the profit value in USD
func profitValue(amount float64) float64 {
	return amount * coin.GetCurrentCoinPrice(util.MDX)
}

// RunCacheClient use cron to poll the vault info at everyday UTC midnight 00:00:00 and every 30
// min to maintain the vault status
func RunCacheClient() {
	initCacher()
	c := cron.New()

	_, err := c.AddFunc("CRON_TZ=UTC @daily", func() {
		// first make sure that the time is 0:0:0 of day
		timestamp := time.Now().Unix()

		stats, err := cacher.VaultStatsFromChain()
		if err != nil {
			logging.Z().Error(
				"[CacheClient]",
				zap.Any("error", err),
			)
			return
		}
		stats.TimeStamp = timestamp
		coll := storage.AccessCollections("vault")
		_, err = coll.InsertOne(context.TODO(), stats)
		if err != nil {
			logging.Z().Error(
				"Insert Db",
				zap.Any("error", err),
				zap.Any("stats", stats),
			)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.AddFunc("CRON_TZ=UTC @every 30m", func() {
		stats, err := cacher.VaultStatsFromChain()
		if err != nil {
			// if fail to get status from the blockchain just give up this update
			return
		}
		stats.TimeStamp = time.Now().Unix()
		cacher.stats = stats
	})
	if err != nil {
		log.Fatal(err)
	}
	c.Start()
}

// StartVaultCache start run the vault cacher, now just start the vault cacher in a subroutine,
func StartVaultCache() {
	go RunCacheClient()
}
