package article

import "github.com/gofiber/fiber/v2"

func getArticle(c *fiber.Ctx)  error{
	return c.SendString("Article")
}