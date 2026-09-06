package core

import (
	"time"

	"github.com/google/uuid"
)

type Vendor struct {
	ID              uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name            string    `gorm:"column:name;not null"`
	Slug            string    `gorm:"column:slug;not null"`
	Description     string    `gorm:"column:description"`
	Email           string    `gorm:"column:email"`
	CreatedOn       time.Time `gorm:"column:created_on;not null"`
	LatestUpdatedOn time.Time `gorm:"column:latest_updated_on;not null"`
	IsActive        bool      `gorm:"column:is_active;not null"`
	IsDeleted       bool      `gorm:"column:is_deleted;not null"`
	Users           []User    `gorm:"foreignKey:VendorID;references:ID"`
}

func (Vendor) TableName() string {
	return "public.vendors"
}
