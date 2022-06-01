package main

import (
	"go-rest-dictionary/config"
	"go-rest-dictionary/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// connect databse while app start up
	config.ConnectDB()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"message": "Welcome to API"})
	})

	// register vocab routes to main app
	routes.VocabRoutes(app)

	app.Listen(":8000")
}
