package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	_id		string `json:"id"`
	Title	string `json:"title"`
	Content	string `json:"content"`
	Author	string `json:"author"`
	CreatedDate	time.Time `json:"created_date"`
}

func GetArticle(c *fiber.Ctx)  error {
	collection, err := getMongoDBCollection(dbName, collectionName)
	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}

	defer cur.Close(context.Background())
	
	cur.All(context.Background(), &results)

	if results == nil {
		return c.SendStatus(404)
		
	}
	// json, _ := json.Marshal(results)
	return c.JSON(results)
}

func CreateArticle(c *fiber.Ctx) error {
	collection, err := getMongoDBCollection(dbName, collectionName)

	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}

	var article Article 

	json.Unmarshal([]byte(c.Body()), &article)

	// test := new(Article)
		
	// fmt.Println(context.Background(), article)
	article.CreatedDate = time.Now()

	res, err := collection.InsertOne(context.Background(), article)

	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}
	json.Marshal(res)

	// response, _ := json.Marshal(res)

	return c.SendString(article.Title)
}

func DeleteArticle (c *fiber.Ctx) error {
	collection, err := getMongoDBCollection(dbName, collectionName)

	if err != nil {
		return c.Status(400).Send([]byte(err.Error()))
	}

	articleID, _ := primitive.ObjectIDFromHex(c.Params("id"))

	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": articleID})

	if err != nil {
		return c.Status(400).Send([]byte(err.Error()))
	}

	response, _ := json.Marshal(res)
	fmt.Println(res)

	return c.Send(response)
}