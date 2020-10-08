package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoConnection es el objeto de conexión a la BD */
var MongoConnection = ConnectionDB()

var clientOpt = options.Client().ApplyURI("mongodb+srv://admin:Holagabster@tuitapp-yivf1.mongodb.net/test?retryWrites=true&w=majority")

/*ConnectionDB es una funcion para conectar nuestra bd */
func ConnectionDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOpt)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println(" Conexion a la base de datos Exitosa ")
	return client
}

/*CheckConnection es una funcion que verifica que haya datos en la conexión a la bd */
func CheckConnection() int {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
