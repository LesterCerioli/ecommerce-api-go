package core

import (
	"time"

	"github.com/google/uuid"
)

type CustomerGroup struct {
	ID              uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name            string     `gorm:"column:name;not null"`
	Description     string     `gorm:"column:description"`
	IsActive        bool       `gorm:"column:is_active;not null"`
	IsDeleted       bool       `gorm:"column:is_deleted;not null"`
	CreatedOn       time.Time  `gorm:"column:created_on;not null"`
	LatestUpdatedOn time.Time  `gorm:"column:latest_updated_on;not null"`
	Users           []CustomerGroupUser `gorm:"foreignKey:CustomerGroupID;references:ID"`
}

func (CustomerGroup) TableName() string {
	return "public.customer_groups"
}
