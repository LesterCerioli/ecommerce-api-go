package contacts

import (
	"time"

	"github.com/google/uuid"
)

type ContactDTO struct {
	ID            uuid.UUID `json:"id"`
	FullName      string    `json:"full_name"`
	PhoneNumber   string    `json:"phone_number"`
	EmailAddress  string    `json:"email_address"`
	Address       string    `json:"address"`
	Content       string    `json:"content"`
	ContactAreaID uuid.UUID `json:"contact_area_id"`
	ContactAreaName string  `json:"contact_area_name,omitempty"`
	IsDeleted     bool      `json:"is_deleted"`
	CreatedOn     time.Time `json:"created_on"`
}
