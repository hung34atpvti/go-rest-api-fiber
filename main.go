package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"rest-api/router"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Static("/", "./public")

	router.RegisterRoutes(app)

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":3000")
}
