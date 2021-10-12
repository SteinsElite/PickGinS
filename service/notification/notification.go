package notification

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/SteinsElite/pickGinS/internal/storage"
)

const (
	QuotaUpdate = "QuotaUpdate"
	Weekly      = "Weekly"
	Activity    = "Activity"
)

const (
	Coll = "notification"
)

type Notification struct {
	Title     string
	Content   string
	Category  string
	TimeStamp int64
}

func PublishNotification(notification Notification) error {
	coll := storage.AccessCollections(Coll)
	_, err := coll.InsertOne(context.TODO(), notification)
	if err != nil {
		log.Println("Fail to write to database due to: ", err)
		return err
	}
	return nil
}

// GetNotification if notification is null string "",
// it means that we should get all the notification
func GetNotification(notification string, page int64, pageSize int64) ([]Notification, int64) {
	coll := storage.AccessCollections(Coll)
	var filter bson.D
	if notification == "" {
		filter = bson.D{}
	} else {
		filter = bson.D{{"category", notification}}
	}

	count, _ := coll.CountDocuments(
		context.TODO(),
		filter,
	)
	opt := options.Find()
	opt.SetLimit(pageSize)
	opt.SetSkip((page - 1) * pageSize)
	opt.SetSort(bson.D{{"timeStamp", -1}})

	cur, err := coll.Find(
		context.TODO(),
		filter,
		opt,
	)
	if err != nil {
		log.Println(err)
	}

	var notifications []Notification
	cur.All(context.TODO(), &notifications)
	defer cur.Close(context.TODO())
	return notifications,count
}
