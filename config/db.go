package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMONGOURI()))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// try connecting to database
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// ping database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB databse")

	return client
}

// client instance
var DB *mongo.Client = ConnectDB()

// getting databse collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("dictionary").Collection(collectionName)

	return collection
}
