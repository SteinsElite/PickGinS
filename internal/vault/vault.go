package vault

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/SteinsElite/pickGinS/internal/coin"
	"github.com/SteinsElite/pickGinS/internal/gateway"
	"github.com/SteinsElite/pickGinS/internal/storage"
	"github.com/SteinsElite/pickGinS/internal/token"
	"github.com/ethereum/go-ethereum/common"
	"github.com/robfig/cron/v3"
)

// maintain the cache of the vault status,provide the interface for other module
// to query
const (
	decimal = 18
)
const (
	Week         = "7D"
	Month        = "1M"
	Year         = "1Y"
	OneDayBefore = -24 * time.Hour
)

type ValuePair struct {
	TimeStamp int64
	Value     float64
}

var vaultWatcher *VaultWatcher

// the vault status on the contract
type VaultStats struct {
	TimeStamp  int64
	CoinAmount map[string]float64
	Profit     float64
}

// use cron to poll the vault info at everyday UTC midnight 00:00:00 and every 30 min to
// maintain the vault status
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
		fmt.Println(err)
	}

	_, err = c.AddFunc("CRON_TZ=UTC @every 30m", func() {
		stats := vaultWatcher.VaultStatsFromChain()
		stats.TimeStamp = time.Now().Unix()
		vaultWatcher.stats = stats
	})
	if err != nil {
		fmt.Println(err)
	}
	c.Start()
}

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

// get the float64 represention of the amount,if necessary, use big.float instead
// (TODO ERIJ)
func (vw *VaultWatcher) tokenAmount(token common.Address) (fAmount float64) {
	tokenState, _ := vw.Instance.TokenState(nil, token)
	tokenVolume := new(big.Int).Sub(tokenState.Max, tokenState.Remain)
	fTokenVolume, _ := new(big.Float).SetString(tokenVolume.String())
	bfAmount := new(big.Float).Quo(fTokenVolume, big.NewFloat(math.Pow10(decimal)))
	fAmount, _ = bfAmount.Float64()
	return
}

func (vw *VaultWatcher) profitAmount() float64 {
	profit, err := vw.Instance.ViewAccumulatedProfit(nil)
	if err != nil {

	}
	fProfit, _ := new(big.Float).SetString(profit.String())
	res, _ := new(big.Float).Quo(fProfit, big.NewFloat(math.Pow10(decimal))).Float64()
	return res
}

// get the latest vault status from blockchain
func (vw *VaultWatcher) VaultStatsFromChain() (stats VaultStats) {
	stats.Profit = vw.profitAmount()
	stats.CoinAmount = make(map[string]float64)
	stats.CoinAmount[token.BTC] = vw.tokenAmount(token.BTCAddr)
	stats.CoinAmount[token.ETH] = vw.tokenAmount(token.ETHAddr)
	stats.CoinAmount[token.USDT] = vw.tokenAmount(token.USDTAddr)
	stats.CoinAmount[token.HT] = vw.tokenAmount(token.HTAddr)
	stats.CoinAmount[token.MDX] = vw.tokenAmount(token.MDXAddr)
	return
}

// the api to be called by other
func queryStartTimeForVolume(phase string) (t time.Time, err error) {
	currentTime := time.Now()
	midnight := midnightOfDay(currentTime)
	switch phase {
	case Week:
		t = midnight.AddDate(0, 0, -7)
	case Month:
		t = midnight.AddDate(0, -1, 0)
	case Year:
		t = midnight.AddDate(-1, 0, 0)
	default:
		err = fmt.Errorf("get the wrong time range")
	}
	return
}

func getQualifiedStatsFromDb(phase string) []VaultStats {
	startTime, err := queryStartTimeForVolume(phase)
	if err != nil {
		log.Println("fail get the start time due to: ", err)
	}
	coll := storage.AccessCollections("vault")
	findOpt := options.Find()
	findOpt.SetSort(bson.D{{"timestamp", 1}})
	cur, err := coll.Find(
		context.Background(),
		bson.D{{"timestamp", bson.D{{"$gte", startTime.Unix()}}}},
		findOpt,
	)
	if err != nil {
		log.Println(err)
	}
	result := []VaultStats{}
	cur.All(context.TODO(), &result)
	defer cur.Close(context.TODO())
	return result
}

func volumeValue(stats VaultStats) float64 {
	var totalValue float64
	for k, amount := range stats.CoinAmount {
		ids, _ := token.TokenIds(k)
		totalValue += amount * coin.GetCurrentCoinPrice(ids)
	}
	return totalValue
}

// return the instant time we shoule lookup in the database
func qulifiedTick(phase string) (tick []int64, err error) {
	current := time.Now()
	midnight := midnightOfDay(current)

	timetick := []int64{}
	switch phase {
	case Week:
		for i := 0; i < 7; i++ {
			tick := midnight.AddDate(0, 0, -1*i-1)
			timetick = append(timetick, tick.Unix())
		}
	case Month:
		for i := 0; i < 30; i++ {
			tick := midnight.AddDate(0, 0, -1*i-1)
			timetick = append(timetick, tick.Unix())
		}
	case Year:
		currentMonth := startOfMonth(current)
		for i := 0; i < 12; i++ {
			tick := currentMonth.AddDate(0, -1*i-1, 0)
			timetick = append(timetick, tick.Unix())
		}
	default:
		err = fmt.Errorf("get the wrong time range")
	}
	return
}

func getQulifiedProfitFromDb(phase string) []ValuePair {
	coll := storage.AccessCollections("vault")
	ticks, _ := qulifiedTick(phase)
	profits := []ValuePair{}
	for _, v := range ticks {
		findOpt := options.Find()
		findOpt.SetSort(bson.D{{"timestamp", 1}})
		findOpt.SetLimit(1)
		cur, err := coll.Find(
			context.Background(),
			bson.D{{"timestamp", bson.D{{"$gte", v}}}},
			findOpt,
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
			profits = append(profits, ValuePair{
				TimeStamp: v,
				Value:     float64(0),
			})
		}
	}
	return profits
}
func profitValue(amount float64) float64 {
	return amount * coin.GetCurrentCoinPrice(token.MDXIds)
}

func midnightOfDay(t time.Time) time.Time {
	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		0, 0, 0, 0,
		time.UTC,
	)
}

func startOfMonth(t time.Time) time.Time {
	return time.Date(
		t.Year(),
		t.Month(),
		0, 0, 0, 0, 0,
		time.UTC,
	)
}
