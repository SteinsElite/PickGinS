package transaction

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/SteinsElite/pickGinS/internal/storage"
)

// The Api Function to interact with the transaction module

func LoadTxFromDb(page, pageSize int64, tag string, address string) ([]TxRecord, error) {
	coll := storage.AccessCollections(txColl)

	opt := options.Find()
	opt.SetLimit(pageSize)
	opt.SetSkip((page - 1) * pageSize)
	opt.SetSort(bson.D{{"timestamp", -1}})

	var filter bson.D
	if tag == "" {
		filter = bson.D{{"user", address}}
	} else {
		filter = bson.D{{"user", address}, bson.E{Key: "tx_type", Value: tag}}
	}
	cur, err := coll.Find(
		context.TODO(),
		filter,
		opt,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(context.TODO())
	var result []TxRecord
	if err = cur.All(context.TODO(), &result); err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

// PollTxInterval will poll the contract periodic to get recent transaction info and persist it
//in the storage. this func will block, so it should start in a standalone goroutine
func PollTxInterval() {
	watcher := InitTxWatcher()
	timeTicker := time.NewTicker(interval * time.Second)
	for {
		latestBlockNumber, err := watcher.RpcClient.Client.BlockNumber(context.TODO())
		if err != nil {
			log.Println("fail to get the block number", err)
		} else {
			tx := watcher.ObtainTxUntil(int64(latestBlockNumber))
			persistRecord(tx)
		}
		<-timeTicker.C
	}
}

func StartTxWatcher() {
	go PollTxInterval()
}
