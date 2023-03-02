package main

import (
	"log"

	"rest-api/config"
	"rest-api/database"
	"rest-api/model"

	"github.com/gofiber/fiber/v2"
)

func main() {

	if err := database.Connect(); err != nil {
		log.Println(err.Error())
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	app.Use(func(c *fiber.Ctx) error {
		log.Println("There is a middleware")
		return c.Next()
	})

	app.Get("/api/v1/organizations/organization", func(c *fiber.Ctx) error {
		data := make([]model.Organization, 0)
		result := database.DB.Find(&data)
		if result.Error != nil {
			c.Status(500).JSON(&fiber.Map{
				"message": "fail",
				"error":   result.Error.Error(),
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": "success",
			"data":    data,
		})
	})

	app.Post("/api/v1/organizations/organization", func(c *fiber.Ctx) error {
		organization := &model.Organization{}
		if err := c.BodyParser(organization); err != nil {
			c.Status(500).JSON(&fiber.Map{
				"message": "fail",
				"error":   err.Error(),
			})
		}
		result := database.DB.Create(&organization)
		if result.Error != nil {
			c.Status(500).JSON(&fiber.Map{
				"message": "fail",
				"error":   result.Error.Error(),
			})
		}
		return c.Status(201).JSON(&fiber.Map{
			"message": "success",
			"data":    organization,
		})
	})

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	if err := app.Listen(":" + config.Config("PORT")); err != nil {
		return
	}

}
