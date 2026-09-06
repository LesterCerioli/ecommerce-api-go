package core

import "github.com/google/uuid"

type EntityDTO struct {
	ID           uuid.UUID `json:"id"`
	Slug         string    `json:"slug"`
	Name         string    `json:"name"`
	EntityID     uuid.UUID `json:"entity_id"`
	EntityTypeID string    `json:"entity_type_id"`
	EntityTypeName string  `json:"entity_type_name,omitempty"`
}
