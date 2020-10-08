package bd

import (
	"context"
	"time"
	"github.com/FelixMH/tuitapp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* ModifyRegister Permite modificar el perfil del usuario. */
func ModifyRegister(u models.Users, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitapp")
	col := db.Collection("users")

	register := make(map[string]interface{})

	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		register["lastName"] = u.LastName
	}
	register["birthday"] = u.Birthday

	if len(u.Email) > 0 {
		register["email"] = u.Email
	}

	if len(u.Password) > 6 && len(u.Password) > 0 {
		u.Password, _ = CryptPassword(u.Password)
	}

	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}

	if len(u.WebSite) > 0 {
		register["webSite"] = u.WebSite
	}

	upd := bson.M{
		"$set":	register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id": bson.M{
			"$eq": objID,
		},
	}

	_, err := col.UpdateOne(ctx, filter, upd)
	if err != nil {
		return false, err
	}
	return true, nil
}