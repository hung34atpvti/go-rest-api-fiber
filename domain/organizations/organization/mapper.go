package organization

import (
	"errors"
	"log"
)

func MapToDTO(model *Organization, dto *DTO) error {
	if model == (&Organization{}) {
		return errors.New("EMPTY_MODEL")
	}
	dto.Id = int(model.ID)
	dto.Name = model.Name
	dto.Description = model.Description
	dto.CreatedAt = model.CreatedAt
	dto.UpdatedAt = model.UpdatedAt
	dto.DeletedAt = model.DeletedAt
	return nil
}

func MapToDTOs(models *[]Organization, dtos *[]DTO) error {
	for _, model := range *models {
		dto := DTO{}
		err := MapToDTO(&model, &dto)
		if err != nil {
			log.Println("EMPTY_MODEL")
		}
		*dtos = append(*dtos, dto)
	}
	return nil
}

func MapToModel(dto *DTO, model *Organization) error {
	if dto == (&DTO{}) {
		return errors.New("EMPTY_DTO")
	}

	model.ID = uint(dto.Id)
	model.Name = dto.Name
	model.Description = dto.Description
	model.CreatedAt = dto.CreatedAt
	model.UpdatedAt = dto.UpdatedAt
	model.DeletedAt = dto.DeletedAt
	return nil
}