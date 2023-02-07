package storage

import (
	"database/sql"
	"rest-api/database/mysql"
)

type MysqlStorage struct {}

func NewMysqlStorage() *MysqlStorage {
	return &MysqlStorage{}
}

func (s *MysqlStorage) Query(query string, params []any)  (*sql.Rows, error) {
	return mysql.DB.Query(query, params...)
}