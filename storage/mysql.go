package storage

import (
	"database/sql"
	"log"
	"rest-api/database/postgres"
	"rest-api/entity"
	"strconv"
)

type MysqlStorage struct {}

func NewMysqlStorage() *MysqlStorage {
	return &MysqlStorage{}
}

func (s *MysqlStorage) FindOrganizations(organization entity.Organization)  (*entity.Organizations, error) {
	query := "SELECT * FROM organization"
	count := 0
	params := make([]string, 0,0)
	if &organization.Id != nil {
		count ++
		if count == 1 {
			query = query + " WHERE "
		} else {
			query = query + ", "
		}
		query = query + "id = $" + strconv.Itoa(count)
		params = append(params, strconv.Itoa(organization.Id))
	}
	if &organization.Description != nil || organization.Description != "" {
		count++
		if count == 1 {
			query = query + " WHERE "
		} else {
			query = query + ", "
		}
		query = query + "description = $" + strconv.Itoa(count)
		params = append(params, organization.Description)
	}
	var err error
	rows, err := postgres.DB.Query(query, params)
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	result := entity.Organizations{}
	for rows.Next() {
		organization := entity.Organization{}
		err := rows.Scan(&organization.Name, &organization.Description)
		// Exit if we get an error
		if err != nil {
			log.Fatal(err)
		}

		// Append Organization to Organizations
		result.Organizations = append(result.Organizations, organization)
	}
	return &result, err
}

func (s *MysqlStorage) FindOrganizationById(organization entity.Organization)  (*entity.Organization, error) {
	return nil, nil
}

func (s *MysqlStorage) CreateOrganization(organization entity.Organization)  (*entity.Organization, error) {
	return nil, nil
}

func (s *MysqlStorage) UpdateOrganization(organization entity.Organization)  (*entity.Organization, error) {
	return nil, nil
}

func (s *MysqlStorage) DeleteOrganization(organization entity.Organization)  (*entity.Organization, error) {
	return nil, nil
}