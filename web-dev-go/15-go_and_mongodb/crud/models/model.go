package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database
var ContextCancel context.CancelFunc
var userCollection *mongo.Collection

func init() {
	ctx, cancelF := context.WithCancel(context.Background())
	ContextCancel = cancelF
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:password@localhost:27017/project1"))
	if err != nil {
		log.Println("ERROR CONNECTING TO MONGODB")
		log.Fatalln(err)
	}
	Database = cl.Database("project1")
	userCollection = Database.Collection("users")
}
