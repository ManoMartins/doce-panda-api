package routes

import (
	"doce-panda/infra/api/handlers"
	"doce-panda/infra/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func CouponRouter(app fiber.Router) {
	app.Get("/coupons", middlewares.EnsureAuthenticated(), handlers.FindAllCoupon())
	app.Get("/coupons/:id", middlewares.EnsureAuthenticated(), handlers.FindByIdCoupon())
	app.Patch("/coupons/:id", middlewares.EnsureAuthenticated(), handlers.UpdateStatusCoupon())
}
