package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/* BackTweet es la estructura que devuelve un tweet. Es decir, como se muestra. */
type BackTweet struct {
	ID		primitive.ObjectID	`bson:"_id" json:"_id, omitempty"`
	UserID	string				`bson:"userID" json:"userID, omitempty"`
	Message	string				`bson:"message" json:"message, omitempty"`
	Date	time.Time			`bson:"date" json:"date, omitempty"`
}
