package catalog

import (
	"time"

	"github.com/google/uuid"

	"ecommerce-api-go/domain/models/core"
)

type Category struct {
	ID              uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name            string     `gorm:"column:name;size:450;not null"`
	Slug            string     `gorm:"column:slug;size:450;not null"`
	MetaTitle       string     `gorm:"column:meta_title;size:450"`
	MetaKeywords    string     `gorm:"column:meta_keywords;size:450"`
	MetaDescription string     `gorm:"column:meta_description"`
	Description     string     `gorm:"column:description"`
	DisplayOrder    int        `gorm:"column:display_order;not null"`
	IsPublished     bool       `gorm:"column:is_published;not null"`
	IncludeInMenu   bool       `gorm:"column:include_in_menu;not null"`
	IsDeleted       bool       `gorm:"column:is_deleted;not null"`
	ParentID        *uuid.UUID `gorm:"column:parent_id;type:uuid"`
	Parent          *Category  `gorm:"foreignKey:ParentID;references:ID"`
	Children        []Category `gorm:"foreignKey:ParentID;references:ID"`
	ThumbnailImageID *uuid.UUID `gorm:"column:thumbnail_image_id;type:uuid"`
	ThumbnailImage  *core.Media `gorm:"foreignKey:ThumbnailImageID;references:ID"`
	CreatedOn       time.Time  `gorm:"column:created_on;not null"`
	LatestUpdatedOn time.Time  `gorm:"column:latest_updated_on;not null"`
}

func (Category) TableName() string {
	return "public.categories"
}
