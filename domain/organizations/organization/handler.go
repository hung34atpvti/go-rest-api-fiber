package organization

import (
	"github.com/gofiber/fiber/v2"
	"rest-api/paging"
	"strconv"
)

func CreateOrganization(c *fiber.Ctx) error {
	organization := new(Organization)
	if err := c.BodyParser(organization); err != nil {
		return err
	}
	data, err := CreateOne(organization)
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

func GetOrganizationById(c *fiber.Ctx) error {
    idStr := c.Params("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": "error",
			"error": err,
		})
	}
	data, err := FindById(id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": "error",
			"error": err,
		})
	}
	if data == (Organization{}) {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Get Successfully",
			"err": "Not Found",
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get Successfully",
		"data": data,
	})
}

func GetOrganizationsAndPagination(c *fiber.Ctx) error {
	pageRequest := new(paging.PageRequest)
	if err := c.QueryParser(pageRequest); err != nil {
		return err
	}

	offset, limit := paging.ConvertToOffsetLimit(pageRequest)

	data, err := FindAllOrganizations(offset, limit)
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




