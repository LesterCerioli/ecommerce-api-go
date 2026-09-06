package catalog

import "github.com/google/uuid"

type ProductOptionCombinationDTO struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	OptionID  uuid.UUID `json:"option_id"`
	OptionName string   `json:"option_name,omitempty"`
	Value     string    `json:"value"`
	SortIndex int       `json:"sort_index"`
}
