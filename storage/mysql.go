package storage

import (
	"database/sql"
	"rest-api/database/mysql"
)

type MysqlStorage struct {
	db *sql.DB
}

func NewMysqlStorage() *MysqlStorage {
	s := new(MysqlStorage)
	s.db = mysql.DB
	return s
}

func (s *MysqlStorage) Query(query string)  (*sql.Rows, error) {
	return s.db.Query(query)
}