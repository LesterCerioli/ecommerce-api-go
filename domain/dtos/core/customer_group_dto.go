package core

import (
	"time"

	"github.com/google/uuid"
)

type CustomerGroupDTO struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	IsActive        bool      `json:"is_active"`
	IsDeleted       bool      `json:"is_deleted"`
	CreatedOn       time.Time `json:"created_on"`
	LatestUpdatedOn time.Time `json:"latest_updated_on"`
}
