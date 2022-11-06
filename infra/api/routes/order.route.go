package routes

import (
	"doce-panda/infra/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func OrderRouter(app fiber.Router) {
	app.Get("/orders/:id", handlers.FindByIdOrder())
	app.Get("/orders", handlers.FindAllOrder())
	app.Post("/orders", handlers.CreateOrder())
}
