package main

import (
	"context"
	"log"
	"mongo-go-example/mongo"
)

func main() {
	ctx := context.Background()
	log.Println("connecting to the db")
	client, err := mongoexample.Example_connect()
	if err != nil {
		log.Fatalf("failed to connect to the db: %v", err)
	}
	insert, err := mongoexample.Example_insert(ctx, client)
	if err != nil {
		return
	}
	log.Println("insert is : ", insert)
	get, err := mongoexample.Example_get_2(client)
	if err != nil {
		return
	}
	log.Println("get is : ", get)
}
