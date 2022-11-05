package api

import (
	"doce-panda/infra/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Server() error {
	app := fiber.New()
	app.Use(cors.New())
	app.Static("/uploads", "./tmp")

	routes.Router(app)

	err := app.Listen(":3333")

	if err != nil {
		return err
	}

	return nil
}
