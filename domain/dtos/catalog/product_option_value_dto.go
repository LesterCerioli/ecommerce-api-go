package catalog

import "github.com/google/uuid"

type ProductOptionValueDTO struct {
	ID          uuid.UUID `json:"id"`
	OptionID    uuid.UUID `json:"option_id"`
	OptionName  string    `json:"option_name,omitempty"`
	ProductID   uuid.UUID `json:"product_id"`
	Value       string    `json:"value"`
	DisplayType string    `json:"display_type"`
	SortIndex   int       `json:"sort_index"`
}
