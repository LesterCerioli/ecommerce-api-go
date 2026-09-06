package catalog

import "github.com/google/uuid"

type ProductAttributeGroupDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
