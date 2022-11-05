package routes

import (
	"doce-panda/infra/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router) {
	app.Get("/users", handlers.FindAllUser())
	app.Get("/users/:id", handlers.FindUser())
	app.Post("/users", handlers.CreateUser())
	app.Post("/users/:id/address", handlers.CreateAddress())
	app.Put("/users/:id", handlers.UpdateUser())
	app.Delete("/users/:id", handlers.DeleteUser())
	app.Delete("/users/:id/address/:addressId", handlers.DeleteAddress())
}
