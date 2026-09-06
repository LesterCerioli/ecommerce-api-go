package tax

import "github.com/google/uuid"

type TaxClassDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
