package bd

import (
	"context"
	"time"
	"fmt"

	"github.com/FelixMH/tuitapp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* ReadAllTweets lee todos los Usuarios registrados en el sistema. */
func ReadUsersAll(ID string, page int64, search string, tipo string) ([]*models.Users, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitapp")
	col := db.Collection("users")

	var results []*models.Users

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println("Error : " + err.Error())
		return results, false
	}

	var OK, include bool

	for cursor.Next(ctx) {
		var s models.Users
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var r models.Relation
		r.UsuarioID = ID
		r.RelationID = s.ID.Hex()

		include = false

		OK, err = SearchRelation(r)
		if tipo == "new" && OK == false {
			include = true
		}
		if tipo == "follow" && OK == true {
			include = true
		}

		if r.RelationID == ID {
			include = false
		}

		if include == true {
			s.Password = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}

	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cursor.Close(ctx)

	return results, true

}
