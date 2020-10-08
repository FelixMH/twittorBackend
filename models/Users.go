package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Users es el modelo de nuestra tabla users */
type Users struct {
	ID			primitive.ObjectID	`bson:"_id,omitempty" json:"id"`
	Name		string				`bson:"name" json:"name,omitempty"`
	LastName	string				`bson:"lastName" json:"lastName,omitempty"`
	Birthday	time.Time			`bson:"birthday" json:"birthday,omitempty"`
	Email		string				`bson:"email" json:"email"`
	Password	string				`bson:"password" json:"password,omitempty"`
	Avatar		string				`bson:"avatar" json:"avatar,omitempty"`
	Banner		string				`bson: "banner" json: "banner, omitempty"`
	WebSite		string				`bson:"webSite" json:"webSite,omitempty"`	
}