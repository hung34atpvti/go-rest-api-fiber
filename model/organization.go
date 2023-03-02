package model

import (
	"gorm.io/gorm"
)

// Organization Model struct
type Organization struct {
	gorm.Model
	Name        string `query:"name"`
	Description string `query:"description"`
}
