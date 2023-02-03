package postgres

import "database/sql"

var DB *sql.DB

func Connect() {
	return
}

func QueryDB(query string) (*sql.Rows, error) {
	return DB.Query(query)
}
