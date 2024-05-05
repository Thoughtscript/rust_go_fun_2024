package domain

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	d "goserver/database"
)

var CollectionName = "examples"

var DatabaseName = "testdatabase"

type Example struct {
	ID		string
	Name	string
	Val		string
}

func GetExamples() []Example {
	var result []Example

	ctx := context.TODO()
	client := d.GetClient()
	collection := client.Database(DatabaseName).Collection(CollectionName)

	// find all without query param - empty bson.D
	cursor, err := collection.Find(ctx, bson.D{});
	if err != nil {
		log.Fatal(err)
	}

	// use cursor only for multi element record set
	for cursor.Next(ctx) {
		var example Example
		cursor.Decode(&example)
		result = append(result, example)
	}

	cursor.Close(ctx)
	return result
}

func GetExample(id string) Example {
	ctx := context.TODO()
	client := d.GetClient()
	collection := client.Database(DatabaseName).Collection(CollectionName)
	filter := bson.D{{"id", id}}

	// find all without query param - empty bson.D
	var result Example
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
    	log.Fatal(err)
	}

	return result
}

func CreateExample(id string, name string, val string) Example {
	ctx := context.TODO()
	client := d.GetClient()
	collection := client.Database(DatabaseName).Collection(CollectionName)
	example := Example{id, name, val}
	
	result, err := collection.InsertOne(ctx, example)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	return example
}

func DeleteExample(id string) Example {
	ctx := context.TODO()
	client := d.GetClient()
	collection := client.Database(DatabaseName).Collection(CollectionName)
	filter := bson.D{{"id", id}}

	var example Example
	if err := collection.FindOne(ctx, filter).Decode(&example); err != nil {
    	log.Fatal(err)
	}
	
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	return example
}

func UpdateExample(id string, name string, val string) Example {
	ctx := context.TODO()
	client := d.GetClient()
	collection := client.Database(DatabaseName).Collection(CollectionName)
	
	example := Example{id, name, val}
	filter := bson.D{{"id", id}}

	result, err := collection.UpdateOne(ctx, filter, example)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	return example
}