package mongoexample

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func Example_insert(ctx context.Context, client *mongo.Client) (*mongo.InsertOneResult, error) {
	db := client.Database("example_data")
	collection := db.Collection("example_collection")

	// Define the document to insert
	document := bson.M{
		"id":        12345,
		"name":      "John Doe",
		"email":     "johndoe@example.com",
		"is_active": true,
		"roles":     []string{"admin", "editor", "viewer"},
		"preferences": bson.M{
			"theme":    "dark",
			"language": "en-US",
			"notifications": bson.M{
				"email": true,
				"sms":   false,
				"push":  true,
			},
		},
		"last_login": "2024-12-18T15:30:00Z",
		"tags":       []string{"random", "json", "example"},
		"metadata":   nil,
	}

	// Insert the document into the collection
	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		log.Fatalf("Failed to insert document: %v", err)
	}

	// Print the inserted document's ID
	fmt.Printf("Inserted document with ID: %v\n", result.InsertedID)
	return result, err
}
