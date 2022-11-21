package routes

import (
	"doce-panda/infra/api/handlers"
	"doce-panda/infra/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router) {
	app.Post("/login", handlers.LoginUser())

	app.Get("/users", middlewares.EnsureAuthenticated(), handlers.FindAllUser())
	app.Get("/users/address", middlewares.EnsureAuthenticated(), handlers.FindAllAddressByUserId())
	app.Get("/users/:id", middlewares.EnsureAuthenticated(), handlers.FindUser())
	app.Post("/users", handlers.CreateUser())
	app.Put("/users/:id", middlewares.EnsureAuthenticated(), handlers.UpdateUser())
	app.Delete("/users/:id", middlewares.EnsureAuthenticated(), handlers.DeleteUser())

	app.Post("/users/address", middlewares.EnsureAuthenticated(), handlers.CreateAddress())
	app.Put("/users/address/:addressId", middlewares.EnsureAuthenticated(), handlers.UpdateAddress())
	app.Delete("/users/address/:addressId", middlewares.EnsureAuthenticated(), handlers.DeleteAddress())
}
