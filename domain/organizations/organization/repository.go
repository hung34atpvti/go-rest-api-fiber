package organization

import (
	"log"
	"rest-api/storage"
)

func GetOrganizations() (Organizations, error) {
	var err error

	err = storage.New("postgres")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := storage.DB.Query("SELECT * FROM organization")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	result := Organizations{}

	for rows.Next() {
		organization := Organization{}
		err := rows.Scan(&organization.Name, &organization.Description)
		// Exit if we get an error
		if err != nil {
			log.Fatal(err)
		}

		// Append Product to Products
		result.Organizations = append(result.Organizations, organization)
	}
	return result, err
}
