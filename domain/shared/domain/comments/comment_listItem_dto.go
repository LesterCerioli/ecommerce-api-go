package comments

import (
	"time"

	"github.com/google/uuid"
)

type CommentListItemDTO struct {
	ID            uuid.UUID     `json:"id"`
	UserID        uuid.UUID     `json:"user_id"`
	CommentText   string        `json:"comment_text"`
	CommenterName string        `json:"commenter_name"`
	Status        CommentStatus `json:"status"`
	CreatedOn     time.Time     `json:"created_on"`
	EntityTypeID  string        `json:"entity_type_id"`
	EntityID      uuid.UUID     `json:"entity_id"`
	EntityName    string        `json:"entity_name"`
	ParentID      *uuid.UUID    `json:"parent_id"`
	EntitySlug    string        `json:"entity_slug"`
}
