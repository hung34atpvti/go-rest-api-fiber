package organization

import (
	"database/sql"
	"log"
	"rest-api/storage"
	"strconv"
)

func FindAllOrganizations(offset int, limit int) ([]Organization, error) {
	err := storage.NewStorage(storage.POSTGRES)
	if err != nil {
		log.Println(err)
	}
	query := "SELECT * FROM organization"
	query = query + " OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(limit)
	params := make([]any, 0, 0)
	rows, err := storage.DB.Query(query, params)
	if err != nil {
		log.Println(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)
	result := Organizations{}

	for rows.Next() {
		organization := Organization{}
		err := rows.Scan(&organization.Id, &organization.Name, &organization.Description)
		// Exit if we get an error
		if err != nil {
			log.Println(err)
		}

		// Append Organization to Organizations
		result.Organizations = append(result.Organizations, organization)
	}
	return result.Organizations, err
}

func CreateOne(organization *Organization) (Organization, error) {
	err := storage.NewStorage(storage.POSTGRES)
	if err != nil {
		log.Println(err)
	}
	query := "INSERT INTO organization (name, description) VALUES ($1, $2)"
	params := make([]any, 0, 0)
	params = append(params, organization.Name)
	params = append(params, organization.Description)
	rows, err := storage.DB.Query(query, params)
	if err != nil {
		log.Println(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)
	result := Organization{}
	for rows.Next() {
		organization := Organization{}
		err := rows.Scan(&organization.Id, &organization.Name, &organization.Description)
		// Exit if we get an error
		if err != nil {
			log.Println(err)
		}

		result = organization
	}
	return result, err
}

func FindById(id int) (Organization, error) {
	err := storage.NewStorage(storage.POSTGRES)
	if err != nil {
		log.Println(err)
	}
	query := "SELECT * FROM organization WHERE id = $1"
	params := make([]any, 0, 0)
	params = append(params, id)
	rows, err := storage.DB.Query(query, params)
	if err != nil {
		log.Println(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)
	result := Organization{}

	for rows.Next() {
		organization := Organization{}
		err := rows.Scan(&organization.Id, &organization.Name, &organization.Description)

		if err != nil {
			log.Println(err)
		}

		// Append Organization to Organizations
		result = organization
	}
	return result, err
}
