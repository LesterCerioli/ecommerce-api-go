package core

import (
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID                   uuid.UUID  `json:"id"`
	UserGUID              uuid.UUID `json:"user_guid"`
	UserName             string    `json:"user_name"`
	NormalizedUserName   string    `json:"normalized_user_name"`
	Email                string    `json:"email"`
	NormalizedEmail      string    `json:"normalized_email"`
	EmailConfirmed       bool      `json:"email_confirmed"`
	PhoneNumber          string    `json:"phone_number"`
	PhoneNumberConfirmed bool      `json:"phone_number_confirmed"`
	TwoFactorEnabled     bool      `json:"two_factor_enabled"`
	LockoutEnabled       bool      `json:"lockout_enabled"`
	AccessFailedCount    int       `json:"access_failed_count"`
	FullName             string    `json:"full_name"`
	VendorID             *uuid.UUID `json:"vendor_id"`
	VendorName           string    `json:"vendor_name,omitempty"`
	IsDeleted            bool      `json:"is_deleted"`
	CreatedOn            time.Time `json:"created_on"`
	LatestUpdatedOn      time.Time `json:"latest_updated_on"`
	Culture              string    `json:"culture"`
}
