package handler

import (
"github.com/gofiber/fiber/v2"
	"log"
	"rest-api/entity"
	"rest-api/storage"
)

func CreateOrganization(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get by Id Successfully",
	})
}

func GetOrganizationById(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get by Id Successfully",
	})
}

func GetOrganizations(c *fiber.Ctx) error {
	o := new(entity.Organization)

	if err := c.QueryParser(o); err != nil {
		log.Fatal(err)
	}
	if err := storage.NewStorage("postgres"); err != nil {
		log.Fatal(err)
	}
	data, err := storage.DB.FindOrganizations(entity.Organization{Id: o.Id, Name: o.Name, Description: o.Description})
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": "error",
			"error": err,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get Successfully",
		"data": data,
	})
}

func UpdateOrganization(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Update Successfully",
	})
}

func RemoveOrganization(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Remove Successfully",
	})
}




