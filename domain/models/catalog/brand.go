package catalog

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	ID              uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name            string    `gorm:"column:name;size:450;not null"`
	Slug            string    `gorm:"column:slug;size:450;not null"`
	Description     string    `gorm:"column:description"`
	IsPublished     bool      `gorm:"column:is_published;not null"`
	IsDeleted       bool      `gorm:"column:is_deleted;not null"`
	CreatedOn       time.Time `gorm:"column:created_on;not null"`
	LatestUpdatedOn time.Time `gorm:"column:latest_updated_on;not null"`
}

func (Brand) TableName() string {
	return "public.brands"
}
