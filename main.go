package main

import (
	"go-rest-dictionary/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// connect databse while app start up
	config.ConnectDB()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"message": "Welcome to API"})
	})

	app.Listen(":8000")
}
