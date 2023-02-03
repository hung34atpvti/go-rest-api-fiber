package team

import "github.com/gofiber/fiber/v2"

func Create(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get by Id Successfully",
	})
}

func FindById(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get by Id Successfully",
	})
}

func Find(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get Successfully",
	})
}

func Update(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Update Successfully",
	})
}

func Remove(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Remove Successfully",
	})
}