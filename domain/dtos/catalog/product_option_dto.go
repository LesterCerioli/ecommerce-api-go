package catalog

import "github.com/google/uuid"

type ProductOptionDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
