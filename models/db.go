package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	connect    *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
}

var secret string
var urldb string

// Init function initializes environment, Cache and DB connections
func InitConnectionToDB(url string, scrt string) (err error) {
	secret = scrt
	urldb = url

	db, err := ConnectToDB()
	if err != nil {
		return
	}
	defer db.Close()

	err = db.connect.Ping(db.ctx, nil)
	return
}

// ConnectionToDB is fuction who create and return DB struct
func ConnectToDB() (db *DB, err error) {
	var ctx context.Context = context.Background()
	var connect *mongo.Client
	connect, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+urldb))
	if err != nil {
		return
	}
	collect := connect.Database("authservice").Collection("session")

	db = &DB{
		connect:    connect,
		collection: collect,
		ctx:        ctx,
	}
	return
}

// Close is method of DB structure who close connection and delete current DB structure
func (db *DB) Close() {
	db.connect.Disconnect(db.ctx)
	db = nil
}
