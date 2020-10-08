package bd

import (
	"context"
	"time"

	"github.com/FelixMH/tuitapp/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* SearchRelation Consulta la relación existente entre dos usuarios. */
func SearchRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitapp")
	col := db.Collection("relation")

	condition := bson.M{
		"userid": t.UsuarioID,
		"relationid": t.RelationID,
	}

	var result models.Relation

	/*fmt.Println(result)*/
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}
	return true, nil

}
