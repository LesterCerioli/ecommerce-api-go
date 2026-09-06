package catalog

import "github.com/google/uuid"

type ProductAttributeValueDTO struct {
	ID          uuid.UUID `json:"id"`
	AttributeID uuid.UUID `json:"attribute_id"`
	AttributeName string  `json:"attribute_name,omitempty"`
	ProductID   uuid.UUID `json:"product_id"`
	Value       string    `json:"value"`
}
