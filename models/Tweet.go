package models

import "time"

/* Tweets Formato o estructura de los datos del Tweet */
type Tweet struct {
	UserID	string		`bson:"userID" json:"userID,omitempty"`
	Message	string		`bson:"message" json:"message,omitempty"`
	Date	time.Time	`bson:"date" json:"date, omitempty"`
}