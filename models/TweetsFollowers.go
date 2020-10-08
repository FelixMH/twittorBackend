package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TweetsFollowers struct {
	ID	primitive.ObjectID `bson:"_id" json:"_id, omitempty"`
	UserID	string `bson:"userId" json:"userId, omitempty"`
	UserRelationID	 string `bson:"userRelationId" json:"userRelationId, omitempty"`
	Tweet struct {
		message string `bson:"message" json:"message, omitempty"`
		Fecha time.Time `bson:"date" json:"date, omitempty"`
		ID	string `bson:"_id", json:"_id, omitempty"`
	}
}