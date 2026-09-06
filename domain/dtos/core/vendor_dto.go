package core

import (
	"time"

	"github.com/google/uuid"
)

type VendorDTO struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Slug            string    `json:"slug"`
	Description     string    `json:"description"`
	Email           string    `json:"email"`
	CreatedOn       time.Time `json:"created_on"`
	LatestUpdatedOn time.Time `json:"latest_updated_on"`
	IsActive        bool      `json:"is_active"`
	IsDeleted       bool      `json:"is_deleted"`
}
