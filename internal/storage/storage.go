package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	timeout     = 2
	mongouri    = "mongodb://192.168.2.200:27017"
	maxpoolsize = 100
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
	// ping will try to select a server until the client's server selection timeout expires.
	opt.SetServerSelectionTimeout(timeout * time.Second)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		// when we fail to connect to the database, we should stop the program
		log.Fatal(err)
	}
	// Call Ping to verify that the deployment is up and the Client was
	// configured successfully. As mentioned in the Ping documentation, this
	// reduces application resiliency as the server may be temporarily
	// unavailable when Ping is called.
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	mgo = client.Database(dbname)
}

func AccessCollections(coll string) *mongo.Collection {
	return mgo.Collection(coll)
}
