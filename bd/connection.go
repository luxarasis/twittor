package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN es el objeto de conneccion a la base de datos */
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://lasis:La1124044454#@cluster0.itc6b9e.mongodb.net/?retryWrites=true&w=majority")

/* ConnectDB es la funcion que me permite conectar a la base de datos */
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Successfull DB connection")

	return client
}

/* CheckConnection es el ping a la base de datos */
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
