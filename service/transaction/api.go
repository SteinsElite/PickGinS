package transaction

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SteinsElite/pickGinS/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// The Api Function to interact with the transaction module

func LoadTxFromDb(page, pageSize int64, tag string, address string) []TxRecord {
	coll := storage.AccessCollections(transaction.txColl)
	findOpt := options.Find()
	findOpt.SetLimit(pageSize)
	findOpt.SetSkip((page - 1) * pageSize)
	findOpt.SetSort(bson.D{{"timestamp", 1}})

	var filter bson.D
	if tag == "" {
		filter = bson.D{{"user", address}}
	} else {
		filter = bson.D{{"user", address}, bson.E{Key: "txtype", Value: tag}}
	}
	cur, err := coll.Find(
		context.Background(),
		filter,
		findOpt,
	)
	if err != nil {
		log.Println(err)
	}
	defer cur.Close(context.Background())
	result := []TxRecord{}
	if err = cur.All(context.Background(), &result); err != nil {
		fmt.Println(err)
	}
	return result
}

// This the watcher of the transaction record, it will poll the contract periodic to
// get recent transaction info and persist it in the storage.
// note: this func will blocking the main goroutine, so it should start in a standlone goroutine
func PollTxInterval() {
	initr := InitRecordObserver()
	timeTicker := time.NewTicker(transaction.interval * time.Second)
	for {
		latestBlockNumber, err := initr.rpcClient().BlockNumber(context.TODO())
		if err != nil {
			log.Println("fail to get the block number", err)
		} else {
			tx := initr.ObtainTxUntil(int64(latestBlockNumber))
			transaction.persistRecord(tx)
		}
		<-timeTicker.C
	}
}
