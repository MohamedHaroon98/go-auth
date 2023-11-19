package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// ConnectDB establishes a connection to MongoDB container
func ConnectDB() *mongo.Client {
	mongo_server := "go-auth-go-auth-mongodb"
	server := fmt.Sprintf("mongodb://%s:27017", mongo_server)

	clientOptions := options.Client().ApplyURI(server)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
