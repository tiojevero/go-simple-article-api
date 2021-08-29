package main

import (
	"github.com/gofiber/fiber/v2"
)

const port string = ":5000" 

func main() {
	app := fiber.New()

	app.Get("/article", article.getArticle)

	app.Listen(port)
}