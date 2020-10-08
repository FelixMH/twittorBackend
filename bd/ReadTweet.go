package bd

import (
	"context"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	"github.com/FelixMH/tuitapp/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* ReadTweet permite realizar la lectura de un tweet. */
func ReadTweet(ID string, page int64) ([]*models.BackTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitapp")
	col := db.Collection("tweet")

	var results []*models.BackTweet

	condition := bson.M{
		"userID": ID,
	}

	options := options2.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{
		{Key: "date", Value: -1},
	})
	options.SetSkip((page - 1)*20)

	cursor, err := col.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var register models.BackTweet

		err := cursor.Decode(&register)
		if err != nil {
			return results, false
		}

		results = append(results, &register)
	}
	return results, true
}

