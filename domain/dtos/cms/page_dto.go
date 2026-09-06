package cms

import (
	"time"

	"github.com/google/uuid"
)

type PageDTO struct {
	ID                uuid.UUID  `json:"id"`
	Name              string     `json:"name"`
	Slug              string     `json:"slug"`
	MetaTitle         string     `json:"meta_title"`
	MetaKeywords      string     `json:"meta_keywords"`
	MetaDescription   string     `json:"meta_description"`
	IsPublished       bool       `json:"is_published"`
	PublishedOn       *time.Time `json:"published_on"`
	IsDeleted         bool       `json:"is_deleted"`
	CreatedByID       uuid.UUID  `json:"created_by_id"`
	CreatedByName     string     `json:"created_by_name,omitempty"`
	CreatedOn         time.Time  `json:"created_on"`
	LatestUpdatedOn   time.Time  `json:"latest_updated_on"`
	LatestUpdatedByID uuid.UUID  `json:"latest_updated_by_id"`
	LatestUpdatedByName string   `json:"latest_updated_by_name,omitempty"`
	Body              string     `json:"body"`
}
