package catalog

import "github.com/google/uuid"

type ProductAttributeDTO struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	GroupID uuid.UUID `json:"group_id"`
	GroupName string  `json:"group_name,omitempty"`
}
