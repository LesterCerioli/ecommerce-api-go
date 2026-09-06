package catalog

import (
	"time"

	"github.com/google/uuid"
)

type BrandDTO struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Slug            string    `json:"slug"`
	Description     string    `json:"description"`
	IsPublished     bool      `json:"is_published"`
	IsDeleted       bool      `json:"is_deleted"`
	CreatedOn       time.Time `json:"created_on"`
	LatestUpdatedOn time.Time `json:"latest_updated_on"`
}
