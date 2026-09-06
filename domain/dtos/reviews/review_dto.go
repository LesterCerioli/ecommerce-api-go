package reviews

import (
	"time"

	"github.com/google/uuid"
)

type ReviewDTO struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	UserName     string    `json:"user_name,omitempty"`
	Title        string    `json:"title"`
	Comment      string    `json:"comment"`
	Rating       int       `json:"rating"`
	ReviewerName string    `json:"reviewer_name"`
	Status       int       `json:"status"`
	CreatedOn    time.Time `json:"created_on"`
	EntityTypeID string    `json:"entity_type_id"`
	EntityID     uuid.UUID `json:"entity_id"`
	EntityName   string    `json:"entity_name,omitempty"`
	EntitySlug   string    `json:"entity_slug,omitempty"`
}
