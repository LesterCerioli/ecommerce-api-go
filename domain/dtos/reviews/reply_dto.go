package reviews

import (
	"time"

	"github.com/google/uuid"
)

type ReplyDTO struct {
	ID          uuid.UUID `json:"id"`
	ReviewID    uuid.UUID `json:"review_id"`
	UserID      uuid.UUID `json:"user_id"`
	UserName    string    `json:"user_name,omitempty"`
	Comment     string    `json:"comment"`
	ReplierName string    `json:"replier_name"`
	Status      int       `json:"status"`
	CreatedOn   time.Time `json:"created_on"`
}
