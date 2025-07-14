package db

import (
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InitDatabase(uri string) *mongo.Client {
	log.Println("MongoDB URI:", uri) // Add this line for debugging

	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(options)
	if err != nil {
		log.Fatalln("Error in connect to database. because: " + err.Error())
	}
	return client
}
