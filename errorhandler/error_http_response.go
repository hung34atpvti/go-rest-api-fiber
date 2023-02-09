package errorhandler

import "github.com/gofiber/fiber/v2"

func InternalServerError(c *fiber.Ctx, err error) error {
	return c.Status(500).JSON(&fiber.Map{
		"message": "Internal Server Error",
		"error": err.Error(),
	})
}

func BadRequest(c *fiber.Ctx, err error) error {
	return c.Status(400).JSON(&fiber.Map{
		"message": "Bad Request",
		"error": err.Error(),
	})
}