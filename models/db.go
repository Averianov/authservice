package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo connection
var collection *mongo.Collection
var ctx = context.Background()

// Secret for jwt interactions
var Secret string

// Init function initializes environment, Cache and DB connections
func Init(urldb string, scrt string) {
	var err error
	Secret = scrt

	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+urldb))
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	collection = client.Database("authservice").Collection("session")
}
