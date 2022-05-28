package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"message": "Welcome to API"})
	})

	app.Listen(":8000")
}
