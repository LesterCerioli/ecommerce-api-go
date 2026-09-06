package core

import "github.com/google/uuid"

type RoleDTO struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	NormalizedName   string    `json:"normalized_name"`
	ConcurrencyStamp string    `json:"concurrency_stamp"`
}
