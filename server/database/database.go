package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBClient *mongo.Client

func ConnectDatabase() {
	log.Println("Database Connecting....")
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))
	ctx := context.TODO()

	// connect
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	DBClient = client

	// ping check connection
	err = DBClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")
}
