package core

import (
	"time"

	"github.com/google/uuid"
)

type Content struct {
	ID                uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name              string     `gorm:"column:name;not null"`
	Slug              string     `gorm:"column:slug;not null"`
	MetaTitle         string     `gorm:"column:meta_title"`
	MetaKeywords      string     `gorm:"column:meta_keywords"`
	MetaDescription   string     `gorm:"column:meta_description"`
	IsPublished       bool       `gorm:"column:is_published;not null"`
	PublishedOn       *time.Time `gorm:"column:published_on"`
	IsDeleted         bool       `gorm:"column:is_deleted;not null"`
	CreatedByID       uuid.UUID  `gorm:"column:created_by_id;type:uuid;not null"`
	CreatedBy         *User      `gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedOn         time.Time  `gorm:"column:created_on;not null"`
	LatestUpdatedOn   time.Time  `gorm:"column:latest_updated_on;not null"`
	LatestUpdatedByID uuid.UUID  `gorm:"column:latest_updated_by_id;type:uuid;not null"`
	LatestUpdatedBy   *User      `gorm:"foreignKey:LatestUpdatedByID;references:ID"`
}
