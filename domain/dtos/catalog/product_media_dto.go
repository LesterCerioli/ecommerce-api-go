package catalog

import "github.com/google/uuid"

type ProductMediaDTO struct {
	ID           uuid.UUID `json:"id"`
	ProductID    uuid.UUID `json:"product_id"`
	MediaID      uuid.UUID `json:"media_id"`
	MediaUrl     string    `json:"media_url,omitempty"`
	DisplayOrder int       `json:"display_order"`
}
