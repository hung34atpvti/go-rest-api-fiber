package storage

import (
	"database/sql"
	"log"
	"rest-api/database/postgres"
	"rest-api/entity"
)

type PostgresStorage struct {}

func NewPostgresStorage() *PostgresStorage {
	return &PostgresStorage{}
}

func (s *PostgresStorage) FindOrganizations(organization entity.Organization)  (*entity.Organizations, error) {
	var query string = "SELECT * FROM organization"
	a := "a"
	log.Fatal(a)
	//count := 0
	//params := make([]any, 0, 0)
	//if organization.Id != 0 {
	//	count ++
	//	if count == 1 {
	//		query = query + " WHERE "
	//	} else {
	//		query = query + " AND "
	//	}
	//	query = query + "id = $" + strconv.Itoa(count)
	//	params = append(params, organization.Id)
	//}
	//if  organization.Description != "" {
	//	count++
	//	if count == 1 {
	//		query = query + " WHERE "
	//	} else {
	//		query = query + " AND "
	//	}
	//	query = query + "name = $" + strconv.Itoa(count)
	//	params = append(params, organization.Name)
	//}
	//if organization.Description != "" {
	//	count++
	//	if count == 1 {
	//		query = query + " WHERE "
	//	} else {
	//		query = query + " AND "
	//	}
	//	query = query + "description = $" + strconv.Itoa(count)
	//	params = append(params, organization.Description)
	//}
	log.Fatal(query)
	//if query[len(query):] == "," {
	//	query = query[0: len(query)-1]
	//}
	var err error
	rows, err := postgres.DB.Query(query, nil)
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
	log.Fatal(result)
	return &result, err
}

func (s *PostgresStorage) FindOrganizationById(organization entity.Organization)  (*entity.Organization, error) {
	return nil, nil
}

func (s *PostgresStorage) CreateOrganization(organization entity.Organization)  (*entity.Organization, error) {
	return nil, nil
}

func (s *PostgresStorage) UpdateOrganization(organization entity.Organization)  (*entity.Organization, error) {
	return nil, nil
}

func (s *PostgresStorage) DeleteOrganization(organization entity.Organization)  (*entity.Organization, error) {
	return nil, nil
}
