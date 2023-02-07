package router

import (
	"github.com/gofiber/fiber/v2"
	"rest-api/handler"
)

func RegisterRoutes(app *fiber.App) {

	apiV1 := app.Group("/api/v1", func (c *fiber.Ctx) error {
		return c.Next()
	})

	//organization route
	apiV1.Get("/organizations/organization", handler.GetOrganizations)
	apiV1.Get("/organizations/organization/:id", handler.GetOrganizationById)
	apiV1.Post("/organizations/organization", handler.CreateOrganization)
	apiV1.Put("/organizations/organization/:id", handler.UpdateOrganization)
	apiV1.Delete("/organizations/organization/:id", handler.RemoveOrganization)
}