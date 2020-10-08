package bd

import (
	"context"
	"time"

	"github.com/FelixMH/tuitapp/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertRegister es la parada final para registrar el user en la db  */
func InsertRegister(u models.Users) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoConnection.Database("tuitapp")
	col := db.Collection("users")

	u.Password, _ = CryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	/* Aqui se obtiene el ID... pero no se manejará así. */
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil

}
