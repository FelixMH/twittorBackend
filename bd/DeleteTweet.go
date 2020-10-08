package bd

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* DeleteTweets borra los tweets del usuario en cuestión */
func DeleteTweets(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitapp")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
		"userID": UserID,
	}

	_, err := col.DeleteOne(ctx, condition)
	return err
}
