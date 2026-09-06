package catalog

import "github.com/google/uuid"

type ProductTemplateDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
