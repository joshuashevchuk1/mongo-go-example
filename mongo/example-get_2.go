package mongoexample

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Example_get_2(client *mongo.Client) ([]byte, error) {
	coll := client.Database("example_data").Collection("example_collection")
	name := "John Doe"
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"name", name}}).
		Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the name %s\n", name)
		return nil, err
	}
	if err != nil {
		panic(err)

	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	return jsonData, err
}
