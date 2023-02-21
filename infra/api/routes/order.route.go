package routes

import (
	"doce-panda/infra/api/handlers"
	"doce-panda/infra/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func OrderRouter(app fiber.Router) {
	app.Get("/orders/me", middlewares.EnsureAuthenticated(), handlers.FindAllOrderByUserId())
	app.Get("/orders/:id", middlewares.EnsureAuthenticated(), handlers.FindByIdOrder())
	app.Get("/orders", middlewares.EnsureAuthenticated(), handlers.FindAllOrder())
	app.Post("/orders", middlewares.EnsureAuthenticated(), handlers.CreateOrder())
	app.Patch("/orders/:id/status", middlewares.EnsureAuthenticated(), handlers.UpdateOrder())
	app.Post("/orders/:id/request-exchange", middlewares.EnsureAuthenticated(), handlers.RequestExchangeOrder())
	app.Patch("/orders/:id/exchange-received", middlewares.EnsureAuthenticated(), handlers.ExchangeReceivedOrder())
	app.Patch("/orders/:id/accept-request-exchange", middlewares.EnsureAuthenticated(), handlers.AcceptRequestExchangeOrder())
	app.Patch("/orders/:id/deny-request-exchange", middlewares.EnsureAuthenticated(), handlers.DenyRequestExchangeOrder())
}
