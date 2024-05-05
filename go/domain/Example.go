package domain

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	d "goserver/database"
)

var CollectionName = "examples"

var DatabaseName = "testdatabase"

// specify the field name, type, and how the field name gets serialized in JSON
type Example struct {
	ID		string `json:id` // not unique
	Name	string `json:name`
	Val		string `json:val`
}

func GetExamples() []Example {
	ctx := context.TODO()
	client := d.GetClient()
	collection := client.Database(DatabaseName).Collection(CollectionName)

	var result []Example = make([]Example, 0)

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
	
	// find all without query param - empty bson.D
	var result Example
	filter := bson.D{{"id", id}}
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
	// _ suppresses/skips the other returned value - use otherwise Go will complain
	_ , err := collection.InsertOne(ctx, example)
	if err != nil {
		log.Fatal(err)
	}

	return example
}

func DeleteExample(id string) Example {
	ctx := context.TODO()
	client := d.GetClient()
	collection := client.Database(DatabaseName).Collection(CollectionName)
	
	var example Example
	filter := bson.D{{"id", id}}
	if err := collection.FindOne(ctx, filter).Decode(&example); err != nil {
    	log.Fatal(err)
	}
	
	_ , err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	return example
}

func UpdateExample(id string, name string, val string) Example {
	ctx := context.TODO()
	client := d.GetClient()
	collection := client.Database(DatabaseName).Collection(CollectionName)
	
	filter := bson.D{{"id", id}}
	// must use the $ notation here
	update := bson.D{{"$set",
		bson.D{
			{ "name", name }, 
			{ "val", val },
		},
	}}

	_ , err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	var example Example
	if err := collection.FindOne(ctx, filter).Decode(&example); err != nil {
    	log.Fatal(err)
	}

	return example
}