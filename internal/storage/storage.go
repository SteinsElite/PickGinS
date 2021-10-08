package storage

import (
"context"
"fmt"
"log"
"time"

"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
)

// package storage provide access to the database, there is a package global
// point mgo represent the connecter,so before use the mgo, call the
// InitDB() first

const (
	timeout      = 10
	mongouri     = "mongodb://localhost:27017"
	dbname       = "pick"
	maxpoolsize  = 100
)

var mgo *mongo.Database

func InitDB(dbname string) {
	connectToDB(mongouri, dbname, maxpoolsize)
}

func connectToDB(uri, dbname string, num uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	opt := options.Client().ApplyURI(uri)
	opt.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		fmt.Println(err)
	}
	mgo = client.Database(dbname)
}

func AccessCollections(coll string) *mongo.Collection {
	if mgo == nil {
		// if mgo is not initialized, we has 2 choice: panic or init it,
		// now we choose init it, so when first access the db collection,
		//we init the db
		InitDB(dbname)
		log.Println("init the db: ", coll)
	}
	return mgo.Collection(coll)
}

