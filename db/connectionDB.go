package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//context = espacio en memoria donde se podrán compartir cosas
//contexto de ejecucion, para setear cosas como timeout
//comunicacion entre ejecuciones
//en resumen  "MANEJO GLOBAL"

// MongoCN - Database Connection instance
var MongoCN = ConnectDB()

//default when push
// var clientOptions = options.Client().ApplyURI("mongodb+srv://username:<password>@murl-tocluster.mongodb.net/test?retryWrites=true&w=majority")
var clientOptions = options.Client().ApplyURI("mongodb+srv://sabal:SabalObzen420#@meganeura-fgfoc.mongodb.net/test?retryWrites=true&w=majority")

//ConnectDB - Database connection
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	//para ver si la conexión
	//de red está disppnible
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión exitosa con Mongo Atlas")
	return client

}

// CheckConnection - Secure the connection
func CheckConnection() int {
	//para ver si la conexión
	//de red está disppnible
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
