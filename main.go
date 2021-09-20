package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

const port string = ":5000" 
const dbName = "go_simple_article_db"
const collectionName = "article"
const mongoURI = "mongodb://localhost:27017"

func Router(app *fiber.App) {
	app.Get("/article/", GetArticle)
	app.Put("/article/:id", UpdateArticle)
	app.Get("/halo", func(c *fiber.Ctx) error {
		return c.SendString("Hallo Bro")
	})
	app.Post("/article", CreateArticle)
	app.Delete("/article/:id", DeleteArticle)
}

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config {
		Prefork: true,
	})
	
	Router(app)

	app.Listen(port)
}