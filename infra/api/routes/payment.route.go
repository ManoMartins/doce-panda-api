package routes

import (
	"doce-panda/infra/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func CreditCardRouter(app fiber.Router) {
	app.Get("/payments/credit-card", handlers.FindAllCreditCard())
	app.Get("/payments/credit-card/:id", handlers.FindByIdCreditCard())
	app.Post("/payments/credit-card", handlers.CreateCreditCard())
	app.Delete("/payments/credit-card/:id", handlers.DeleteCreditCard())
}
