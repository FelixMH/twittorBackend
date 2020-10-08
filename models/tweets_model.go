package models

/* Tweets captura los tweets que llegan del body. */
type Tweets struct {
	Message		string	`bson:"message" json:"message"`
}
