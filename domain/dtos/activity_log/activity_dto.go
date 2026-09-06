package activity_log

import (
	"time"

	"github.com/google/uuid"
)

type ActivityDTO struct {
	ID             uuid.UUID `json:"id"`
	ActivityTypeID uuid.UUID `json:"activity_type_id"`
	ActivityTypeName string  `json:"activity_type_name,omitempty"`
	UserID         uuid.UUID `json:"user_id"`
	UserName       string    `json:"user_name,omitempty"`
	CreatedOn      time.Time `json:"created_on"`
	EntityID       uuid.UUID `json:"entity_id"`
	EntityTypeID   string    `json:"entity_type_id"`
}
