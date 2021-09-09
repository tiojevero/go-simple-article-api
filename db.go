package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


func GetMongoDBConnection() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Println(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	
	if err != nil {
		fmt.Println(err)
	}

	return client, nil
} 

func getMongoDBCollection(DBName string, CollectionName string) (*mongo.Collection, error) {
	client, err := GetMongoDBConnection()

	if err != nil {
		fmt.Println(err)
	}

	collection := client.Database(DBName).Collection(CollectionName)

	return collection, nil
}