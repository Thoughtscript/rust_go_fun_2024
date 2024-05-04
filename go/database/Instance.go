package database

import (
	"context"
	"sync"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client


// Define mutex
var dbLock = &sync.Mutex{}

func GetClient() *mongo.Client {
	dbLock.Lock()
	defer dbLock.Unlock()

	if client != nil {
		return client
	}

	var err error
	
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))
	log.Println("Database connection opened...")
	if err != nil {
		log.Fatal(err)
	}
	
	return client
}
