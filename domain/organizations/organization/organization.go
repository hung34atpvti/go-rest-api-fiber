package organization

import (
	"time"

	"gorm.io/gorm"
)

// Organization Model struct
type Organization struct {
	gorm.Model
	Name        string
	Description string
}

// DTO struct
type DTO struct {
	Id          int            `json:"id" query:"id"`
	Name        string         `json:"name" query:"name"`
	Description string         `json:"description" query:"description"`
	CreatedAt   time.Time      `json:"createdAt" query:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt" query:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" query:"deletedAt"`
}
