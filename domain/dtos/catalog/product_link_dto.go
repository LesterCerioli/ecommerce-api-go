package catalog

import "github.com/google/uuid"

type ProductLinkDTO struct {
	ID              uuid.UUID `json:"id"`
	ProductID       uuid.UUID `json:"product_id"`
	ProductName     string    `json:"product_name,omitempty"`
	LinkedProductID uuid.UUID `json:"linked_product_id"`
	LinkedProductName string  `json:"linked_product_name,omitempty"`
	LinkType        int       `json:"link_type"`
}
