package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/FelixMH/tuitapp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
/* SearchProfile busca un perfil en la base de datos  */
func SearchProfile(ID string) (models.Users, error) {
	ctx , cancel := context.WithTimeout(context.Background(), time.Second * 15)
	defer cancel()

	db := MongoConnection.Database("tuitapp")
	col := db.Collection("users")

	var profile models.Users
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":	objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Perfil no encontrado..."+err.Error())
		return profile, err
	}
	return profile, nil


}