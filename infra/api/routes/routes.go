package routes

import "github.com/gofiber/fiber/v2"

func Router(app fiber.Router) {
	ProductRouter(app)
	UserRouter(app)
	OrderRouter(app)
	CreditCardRouter(app)
}
