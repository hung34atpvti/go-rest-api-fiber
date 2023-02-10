package router

import (
	"rest-api/domain/organizations/organization"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	apiV1 := app.Group("/api/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})

	//organization route
	apiV1.Get("/organizations/organization", organization.GetOrganizationsAndPaginationHandler)
	apiV1.Get("/organizations/organization/:id", organization.GetOrganizationByIdHandler)
	apiV1.Post("/organizations/organization", organization.CreateOrganizationHandler)
	apiV1.Put("/organizations/organization/:id", organization.UpdateOrganizationHandler)
	apiV1.Delete("/organizations/organization/:id", organization.RemoveOrganizationHandler)
}
