package domain

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	d "goserver/database"
)

type Example struct {
	ID		string
	Name	string
	Val		string
}

func GetExamples() []Example {
	var result []Example

	ctx := context.TODO()
	client := d.GetClient()
	collection := client.Database("testdatabase").Collection("examples")

	// find all without query param - empty bson.D
	cursor, err := collection.Find(ctx, bson.D{});
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var entry bson.M
		cursor.Decode(&entry)

		example := Example{
			ID:		"1",
			Name:	"1",
			Val:	"1",
		}

		log.Println(example)
		result = append(result, example)
	}

	return result
}