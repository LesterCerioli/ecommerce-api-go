package news

import "github.com/google/uuid"

type NewsCategoryDTO struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Slug            string    `json:"slug"`
	MetaTitle       string    `json:"meta_title"`
	MetaKeywords    string    `json:"meta_keywords"`
	MetaDescription string    `json:"meta_description"`
	Description     string    `json:"description"`
	DisplayOrder    int       `json:"display_order"`
	IsPublished     bool      `json:"is_published"`
	IsDeleted       bool      `json:"is_deleted"`
}
