package catalog

import "github.com/google/uuid"

type ProductCategoryDTO struct {
	ID                uuid.UUID `json:"id"`
	IsFeaturedProduct bool      `json:"is_featured_product"`
	DisplayOrder      int       `json:"display_order"`
	CategoryID        uuid.UUID `json:"category_id"`
	CategoryName      string    `json:"category_name,omitempty"`
	CategorySlug      string    `json:"category_slug,omitempty"`
	ProductID         uuid.UUID `json:"product_id"`
	ProductName       string    `json:"product_name,omitempty"`
	ProductSlug       string    `json:"product_slug,omitempty"`
}
