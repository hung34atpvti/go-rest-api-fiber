package router

import (
	"github.com/gofiber/fiber/v2"
	"rest-api/domain/organizations/organization"
	"rest-api/domain/organizations/team"
)

func RegisterRoutes(app *fiber.App) {

	apiV1 := app.Group("/api/v1", func (c *fiber.Ctx) error {
		return c.Next()
	})

	//organization route
	apiV1.Get("/organizations/organization", organization.Find)
	apiV1.Get("/organizations/organization/:id", organization.FindById)
	apiV1.Post("/organizations/organization", organization.Create)
	apiV1.Put("/organizations/organization/:id", organization.Update)
	apiV1.Delete("/organizations/organization/:id", organization.Remove)

	//team route
	apiV1.Get("/organizations/team", team.Find)
	apiV1.Get("/organizations/team/:id", team.FindById)
	apiV1.Post("/organizations/team", team.Create)
	apiV1.Put("/organizations/team/:id", team.Update)
	apiV1.Delete("/organizations/team/:id", team.Remove)
}