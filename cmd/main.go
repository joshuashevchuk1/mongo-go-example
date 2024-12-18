package main

import (
	"log"
	"mongo-go-example/mongo"
)

func main() {
	log.Println("connecting to the db")
	client, err := mongoexample.Example_connect()
	if err != nil {
		log.Fatalf("failed to connect to the db: %v", err)
	}
	get, err := mongoexample.Example_get_2(client)
	if err != nil {
		return
	}
	log.Println("get is : ", get)
}
