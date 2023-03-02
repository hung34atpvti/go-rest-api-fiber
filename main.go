package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	app.Use(func(c *fiber.Ctx) error {
		log.Println("There is a middleware")
		return c.Next()
	})

	app.Get("/api/v1/organizations/organization", func(c *fiber.Ctx) error {
		return c.SendString("Hello from organization")
	})

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	if err := app.Listen(":3000"); err != nil {
		return
	}

}
