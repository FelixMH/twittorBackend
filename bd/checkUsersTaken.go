package bd

import (
	"context"
	"time"

	"github.com/FelixMH/tuitapp/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckUsersTaken Revisa si un email ya está registrado en la base de datos.  */
func CheckUsersTaken(email string) (models.Users, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitapp")
	col := db.Collection("users")

	condicion := bson.M{"email": email}

	var resultado models.Users

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
