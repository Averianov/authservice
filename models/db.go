package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo connection
var collection *mongo.Collection
var ctx = context.Background()

// Secret for jwt interactions
var Secret string

// Init function initializes environment, Cache and DB connections
func InitDB(urldb string, scrt string) (err error) {
	Secret = scrt

	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+urldb))
	if err = client.Ping(ctx, nil); err != nil {
		return
	}

	collection = client.Database("authservice").Collection("session")
	return
}
