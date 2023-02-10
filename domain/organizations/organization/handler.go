package organization

import (
	"rest-api/errorhandler"
	"rest-api/paging"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateOrganizationHandler(c *fiber.Ctx) error {
	reqDto := &DTO{}
	if err := c.BodyParser(reqDto); err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	reqOrganization := &Organization{}
	if err := MapToModel(reqDto, reqOrganization); err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	organizationModel, err := CreateOne(reqOrganization)
	if err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	organizationDto := &DTO{}
	if err := MapToDTO(organizationModel, organizationDto); err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	return c.Status(201).JSON(&fiber.Map{
		"message": "Organization Created",
		"data":    organizationDto,
	})
}

func GetOrganizationByIdHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	model, err := FindById(id)
	if err != nil {
		if err.Error() == "record not found" {
			return errorhandler.BadRequest(c, err)
		}
		return errorhandler.InternalServerError(c, err)
	}

	organizationDto := &DTO{}
	if err := MapToDTO(model, organizationDto); err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"message": "Get Organization",
		"data":    organizationDto,
	})
}

func GetOrganizationsAndPaginationHandler(c *fiber.Ctx) error {
	pageRequest := &paging.PageRequest{}
	if err := c.QueryParser(pageRequest); err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	offset, limit := paging.ConvertToOffsetLimit(pageRequest)

	organizationModels, err := FindAllOrganizations(offset, limit)
	if err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	organizationDtos := make([]DTO, 0)
	if err := MapToDTOs(organizationModels, &organizationDtos); err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"message": "Get Organizations",
		"data":    organizationDtos,
	})
}

func UpdateOrganizationHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	reqDto := &DTO{}
	if err := c.BodyParser(reqDto); err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	reqModel := &Organization{}
	if err := MapToModel(reqDto, reqModel); err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	organizationModel, err := UpdateById(id, reqModel)
	if err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	organizationDto := &DTO{}
	if err := MapToDTO(organizationModel, organizationDto); err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"message": "Organization Updated",
		"data":    organizationDto,
	})
}

func RemoveOrganizationHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	model, err := DeleteById(id)
	if err != nil {
		if err.Error() == "record not found" {
			return errorhandler.BadRequest(c, err)
		}
		return errorhandler.InternalServerError(c, err)
	}

	organizationDto := &DTO{}
	if err := MapToDTO(model, organizationDto); err != nil {
		return errorhandler.InternalServerError(c, err)
	}

	organizationDto.Id = id

	return c.Status(200).JSON(&fiber.Map{
		"message": "Organization Removed",
		"data":    organizationDto,
	})
}
