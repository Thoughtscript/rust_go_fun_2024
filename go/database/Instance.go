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

	var connectionURI = "mongodb://testuser:testpass@mongodb:27017/testdatabase?authSource=testdatabase"
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionURI))
	log.Println("Database connection opened...")
	if err != nil {
		log.Fatal(err)
	}
	client = c
	return client
}
