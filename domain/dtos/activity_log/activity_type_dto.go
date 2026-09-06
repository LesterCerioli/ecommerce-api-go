package activity_log

import "github.com/google/uuid"

type ActivityTypeDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
