package bd

import (
	"context"
	"time"
	"github.com/FelixMH/tuitapp/models"
)

/* InsertRelation guarda la relacion entre los usuarios en una coleccion relation de mongo. */
func InsertRelation (t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitapp")
	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}