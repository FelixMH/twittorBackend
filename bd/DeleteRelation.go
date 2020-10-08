package bd

import (
	"context"
	"time"
	"github.com/FelixMH/tuitapp/models"
)

/* DeleteRelation borra la relacion de Mongo */
func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitapp")
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}

