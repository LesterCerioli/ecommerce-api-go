package reviews

import (
	"time"

	"github.com/google/uuid"
)

type ReplyListItemDTO struct {
	ID          uuid.UUID `json:"id"`
	Comment     string    `json:"comment"`
	ReplierName string    `json:"replier_name"`
	Status      int       `json:"status"`
	CreatedOn   time.Time `json:"created_on"`
	ReviewTitle string    `json:"review_title"`
	EntityName  string    `json:"entity_name"`
	EntitySlug  string    `json:"entity_slug"`
}
