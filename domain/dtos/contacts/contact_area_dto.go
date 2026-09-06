package contacts

import "github.com/google/uuid"

type ContactAreaDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsDeleted bool      `json:"is_deleted"`
}
