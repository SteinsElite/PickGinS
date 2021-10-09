package notification

import (
	"context"
	"github.com/SteinsElite/pickGinS/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
	TimeStamp uint64
}

func PublishNotification(notification Notification) {
	coll := storage.AccessCollections(Coll)
	_, err := coll.InsertOne(context.TODO(), notification)
	if err != nil {
		log.Println("Fail to write to database due to: ", err)
	}
}

// GetNotification if noti is null string "", it mean that we should get all the notification
func GetNotification(notification string, page int64, pageSize int64) []Notification {
	coll := storage.AccessCollections(Coll)

	opt := options.Find()
	opt.SetLimit(pageSize)
	opt.SetSkip((page - 1) * pageSize)
	opt.SetSort(bson.D{{"TimeStamp", -1}})

	var filter bson.D
	if notification == "" {
		filter = bson.D{}
	} else {
		filter = bson.D{{"Category", notification}}
	}
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
	return notifications
}
