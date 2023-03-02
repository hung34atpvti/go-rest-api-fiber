package main

import (
	"log"
	"rest-api/config"
	"rest-api/database/postgres"
	"rest-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	if err := postgres.Connect(); err != nil {
		log.Println(err.Error())
	}

	app := fiber.New()

	app.Use(logger.New())

	app.Static("/", "./public")

	router.RegisterRoutes(app)

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	if err := app.Listen(":" + config.Config("PORT")); err != nil {
		return
	}

}
