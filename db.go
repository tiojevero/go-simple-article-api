package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type MongoInstance struct {
	Client	*mongo.Client
	Db		*mongo.Database
}

var mg MongoInstance

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db: db,
	}

	return nil
}



// func GetMongoDBConnection() (*mongo.Client, error) {
// 	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = client.Ping(context.Background(), readpref.Primary())
	
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return client, nil
// } 

// func getMongoDBCollection(DBName string, CollectionName string) (*mongo.Collection, error) {
// 	client, err := GetMongoDBConnection()

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	collection := client.Database(DBName).Collection(CollectionName)

// 	return collection, nil
// }