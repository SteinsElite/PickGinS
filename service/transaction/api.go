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

// LoadTxFromDb query transaction in db and return with valid tx record and total count of
// specific record
func LoadTxFromDb(page, pageSize int64, tag string, address string) ([]TxRecord, int64, error) {
	coll := storage.AccessCollections(txColl)

	var filter bson.D
	if tag == "" {
		filter = bson.D{{"user", address}}
	} else {
		filter = bson.D{{"user", address}, {"tx_type", tag}}
	}

	count, err := coll.CountDocuments(
		context.TODO(),
		filter,
	)
	if err != nil {
		return nil, 0, err
	}

	var result []TxRecord
	opt := options.Find()
	opt.SetLimit(pageSize)
	opt.SetSkip((page - 1) * pageSize)
	opt.SetSort(bson.D{{"timestamp", -1}})
	cur, err := coll.Find(
		context.TODO(),
		filter,
		opt,
	)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(context.TODO())

	if err = cur.All(context.TODO(), &result); err != nil {
		return nil, 0, err
	}
	return result, count, nil
}

// PollTxInterval will poll the contract periodic to get recent transaction info and persist it
// in the storage. this func will block, so it should start in a standalone goroutine
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
