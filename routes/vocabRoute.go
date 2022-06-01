package routes

import (
	"go-rest-dictionary/controller"

	"github.com/gofiber/fiber/v2"
)

func VocabRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/vocab", controller.StoreVocab)
	api.Get("/vocab", controller.IndexVocab)
	api.Get("/vocab/:word", controller.Translate)
	api.Put("/vocab/:id", controller.UpdateVocab)
	api.Delete("/vocab/:id", controller.DeleteVocab)
}
