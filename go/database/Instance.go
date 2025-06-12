package database

import (
	"context"
	"log"
	"sync"

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

	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://testuser:testpass@mongodb:27017/testdatabase"))
	log.Println("Database connection opened...")
	if err != nil {
		log.Fatal(err)
	}
	client = c
	return client
}
