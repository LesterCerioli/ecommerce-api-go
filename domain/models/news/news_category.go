package news

import "github.com/google/uuid"

type NewsCategory struct {
	ID              uuid.UUID          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name            string             `gorm:"column:name;size:450;not null"`
	Slug            string             `gorm:"column:slug;size:450;not null"`
	MetaTitle       string             `gorm:"column:meta_title;size:450"`
	MetaKeywords    string             `gorm:"column:meta_keywords;size:450"`
	MetaDescription string             `gorm:"column:meta_description"`
	Description     string             `gorm:"column:description"`
	DisplayOrder    int                `gorm:"column:display_order;not null"`
	IsPublished     bool               `gorm:"column:is_published;not null"`
	IsDeleted       bool               `gorm:"column:is_deleted;not null"`
	NewsItems       []NewsItemCategory `gorm:"foreignKey:CategoryID;references:ID"`
}

func (NewsCategory) TableName() string {
	return "public.news_categories"
}
