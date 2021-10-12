package vault

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/SteinsElite/pickGinS/internal/gateway"
	"github.com/SteinsElite/pickGinS/internal/storage"
	"github.com/SteinsElite/pickGinS/service/coin"
	"github.com/SteinsElite/pickGinS/util"
)

// vault.go maintain the cache of the vault status,provide the interface for other module
// to query

const (
	Week  = "7D"
	Month = "1M"
	Year  = "1Y"
)

// vault watcher will store the latest vault status since last update which
// will be queried by other parts
var vaultWatcher *VaultWatcher

type ValuePair struct {
	TimeStamp int64
	Value     float64
}

// VaultStats the vault status on the contract
type VaultStats struct {
	TimeStamp  int64
	CoinAmount map[string]float64
	Profit     float64 // total profit util now
}

// set up the vault watcher to provide service for querying
func initVaultWatcher() {
	vaultWatcher = &VaultWatcher{
		RpcClient: *gateway.GetRpcClient(),
	}
	vaultWatcher.stats = vaultWatcher.VaultStatsFromChain()
	vaultWatcher.stats.TimeStamp = time.Now().Unix()
}

type VaultWatcher struct {
	gateway.RpcClient
	stats VaultStats
}

// get the float64 representation of the amount,if necessary, use big.float instead
func (vw *VaultWatcher) tokenAmount(token common.Address) (fAmount float64) {
	tokenState, _ := vw.Instance.TokenState(nil, token)
	tokenVolume := new(big.Int).Sub(tokenState.Max, tokenState.Remain)
	fAmount = util.Amount2Float(tokenVolume)
	return
}

func (vw *VaultWatcher) profitAmount() (fProfit float64) {
	profit, _ := vw.Instance.ViewAccumulatedProfit(nil)
	fProfit = util.Amount2Float(profit)
	return
}

// VaultStatsFromChain get the latest vault status from blockchain,
func (vw *VaultWatcher) VaultStatsFromChain() (stats VaultStats) {
	stats.Profit = vw.profitAmount()
	stats.CoinAmount = make(map[string]float64)
	stats.CoinAmount[util.BTC] = vw.tokenAmount(util.BTCAddr)
	stats.CoinAmount[util.ETH] = vw.tokenAmount(util.ETHAddr)
	stats.CoinAmount[util.USDT] = vw.tokenAmount(util.USDTAddr)
	stats.CoinAmount[util.HT] = vw.tokenAmount(util.HTAddr)
	stats.CoinAmount[util.MDX] = vw.tokenAmount(util.MDXAddr)
	return
}

// get the qualified status start since the start time
func getQualifiedStatsFromDb(startTime int64) []VaultStats {
	coll := storage.AccessCollections("vault")
	opt := options.Find()
	opt.SetSort(bson.D{{"timestamp", 1}})
	cur, err := coll.Find(
		context.TODO(),
		bson.D{{"timestamp", bson.D{{"$gte", startTime}}}},
		opt,
	)
	if err != nil {
		log.Println(err)
	}
	var result []VaultStats
	cur.All(context.TODO(), &result)
	defer cur.Close(context.TODO())
	return result
}

func getQualifiedProfitFromDb(ticks []int64) []ValuePair {
	coll := storage.AccessCollections("vault")
	var profits []ValuePair
	for _, v := range ticks {
		opt := options.Find()
		opt.SetSort(bson.D{{"timestamp", -1}})
		opt.SetLimit(1)
		// find record that timestamp is less than or equal to the tick time
		cur, err := coll.Find(
			context.TODO(),
			bson.D{{"timestamp", bson.D{{"$lte", v}}}},
			opt,
		)
		if err != nil {
			log.Println(err)
		}
		if cur.Next(context.TODO()) {
			var stats VaultStats
			err := cur.Decode(&stats)
			if err != nil {
				log.Println(err)
			}
			profits = append(profits, ValuePair{
				TimeStamp: v,
				Value:     stats.Profit,
			})
		} else {
			// if we can't get the stats at the timestamp: we don't start the vault watcher since
			// that time, so just set it to 0
			profits = append(profits, ValuePair{
				TimeStamp: v,
				Value:     0.0,
			})
		}
	}
	return profits
}

// calculate the total volume value in USD of the specific stats
func volumeValue(stats VaultStats) float64 {
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

// RunVaultWatcher use cron to poll the vault info at everyday UTC midnight 00:00:00 and every 30
// min to maintain the vault status
func RunVaultWatcher() {
	initVaultWatcher()
	c := cron.New()

	_, err := c.AddFunc("CRON_TZ=UTC @daily", func() {
		timestamp := time.Now().Unix()
		vaultStats := vaultWatcher.VaultStatsFromChain()
		vaultStats.TimeStamp = timestamp
		coll := storage.AccessCollections("vault")
		_, err := coll.InsertOne(context.TODO(), vaultStats)
		if err != nil {
			fmt.Println(err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.AddFunc("CRON_TZ=UTC @every 30m", func() {
		stats := vaultWatcher.VaultStatsFromChain()
		stats.TimeStamp = time.Now().Unix()
		vaultWatcher.stats = stats
	})
	if err != nil {
		log.Fatal(err)
	}
	c.Start()
}

// StartVaultWatcher start run the vault watcher, now just start the vault watcher in a subroutine,
// (TODO(ERIJ)) try to deal with the error in the StartVaultWacther
func StartVaultWatcher() {
	go RunVaultWatcher()
}
