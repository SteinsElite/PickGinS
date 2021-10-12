package router

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/SteinsElite/pickGinS/internal/storage"
	"github.com/SteinsElite/pickGinS/util"
)

// auth.go point out that who is allowed to publish the notification

const (
	RootAddr = "0x397a9e1719113Cd68ba79d59e1e1988C669cA7F3"
	RootWord = "PICK"
	coll = "admin"
)

type AuthAccount struct {
	Address  common.Address
	WordHash []byte
}

func IsAuth(addr string, sig string) bool {
	sigByte := []byte(sig)
	dataHash := getAuthWord(addr)
	if dataHash == nil {
		// the account is not register in the db
		return false
	}
	if util.VerifySignature(addr, dataHash, sigByte) {
		// test that the sig is sign by the address
		return true
	}
	return false
}

// InitAuth store the first root account in the db to publish the notification
func InitAuth() {
	coll := storage.AccessCollections(coll)
	count, _ := coll.CountDocuments(
		context.TODO(),
		bson.D{},
	)
	if count > 0 {
		return
	}
	_, err := coll.InsertOne(context.TODO(), AuthAccount{
		Address:  common.HexToAddress(RootAddr),
		WordHash: crypto.Keccak256Hash([]byte(RootWord)).Bytes(),
	})
	if err != nil {
		log.Println("Fail writing to the database: ", err)
	}
}

func getAuthWord(address string) []byte {
	addr := common.HexToAddress(address)
	coll := storage.AccessCollections(coll)

	cur, err := coll.Find(
		context.TODO(),
		bson.D{{"address", addr}},
	)
	if err != nil {
		log.Println(err)
	}
	defer cur.Close(context.TODO())

	if cur.Next(context.TODO()) {
		var authAccount AuthAccount
		err := cur.Decode(&authAccount)
		if err != nil {
			log.Println(err)
		}
		return authAccount.WordHash
	}
	return nil
}

func SetNewPublisher(address string, word string) {
	coll := storage.AccessCollections(coll)
	_, err := coll.InsertOne(context.TODO(), AuthAccount{
		Address:  common.HexToAddress(address),
		WordHash: crypto.Keccak256Hash([]byte(word)).Bytes(),
	})
	if err != nil {
		log.Println("Fail writing to the database: ", err)
	}
}
