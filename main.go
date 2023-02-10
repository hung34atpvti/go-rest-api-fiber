package main

import (
	"log"
	"rest-api/database/postgresdb"
	"rest-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	if err := postgresdb.Connect(); err != nil {
		log.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	app.Static("/", "./public")

	router.RegisterRoutes(app)

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	if err := app.Listen(":3000"); err != nil {
		return
	}

}
