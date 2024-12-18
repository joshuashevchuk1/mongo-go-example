// Connects to MongoDB and sets a Stable API version
package main

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Replace the placeholder with your Atlas connection string
const uri = "<connection string>"

func main() {

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	log.Println("Connecting to MongoDB", serverAPI)

}
