package main

import (
	"github.com/gofiber/fiber/v2"
)

const port string = ":5000" 

func Router(app *fiber.App) {
	app.Get("/article", article.getArticle)
}

func main() {
	app := fiber.New()
	Router(app)

	app.Listen(port)
}