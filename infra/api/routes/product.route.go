package routes

import (
	"doce-panda/infra/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func ProductRouter(app fiber.Router) {
	app.Get("/products", handlers.FindAllProduct())
	app.Get("/products/:id", handlers.FindProduct())
	app.Post("/products", handlers.CreateProduct())
	app.Put("/products/:id", handlers.UpdateProduct())
	app.Delete("/products/:id", handlers.DestroyProduct())
	app.Patch("/products/:id/upload", handlers.UploadProduct())
	app.Patch("/products/:id/enable", handlers.EnableProduct())
	app.Patch("/products/:id/disable", handlers.DisableProduct())

	app.Post("/categories", handlers.CreateCategory())
}
