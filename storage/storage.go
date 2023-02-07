package storage

import (
	"rest-api/entity"
)

type Storage interface {
	FindOrganizations(entity.Organization) (*entity.Organizations, error)
	FindOrganizationById(entity.Organization) (*entity.Organization, error)
	CreateOrganization(entity.Organization) (*entity.Organization, error)
	UpdateOrganization(entity.Organization) (*entity.Organization, error)
	DeleteOrganization(entity.Organization) (*entity.Organization, error)
}

var DB Storage

func NewStorage(dbType string) error {
	if dbType == "postgres" {
		DB = NewPostgresStorage()
	}
	if dbType == "mysql" {
		DB = NewMysqlStorage()
	}
	return nil
}
