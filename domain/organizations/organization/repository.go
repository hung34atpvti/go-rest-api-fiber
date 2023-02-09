package organization

import (
	"rest-api/database/postgresdb"
)

func FindAllOrganizations(offset int, limit int) (*[]Organization, error) {
	data := make([]Organization, 0, 0)
	result := postgresdb.DB.Find(&data).Limit(limit).Offset(offset)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func CreateOne(organization *Organization) (*Organization, error) {
	result := postgresdb.DB.Create(&organization)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return organization, nil
}

func FindById(id int) (*Organization, error) {
	organization := Organization{}
	result := postgresdb.DB.First(&organization, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &organization, nil
}

func UpdateById(id int, organization *Organization) (*Organization, error) {
	_, err := FindById(id)
	if err != nil {
		return nil, err
	}
	organization.ID = uint(id)
	result := postgresdb.DB.Save(&organization)
	if result.Error != nil {
		return nil, result.Error
	}
	return organization, nil
}
