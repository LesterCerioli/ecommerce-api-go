package core

import (
	"time"

	"github.com/google/uuid"
)

type UserAddressDTO struct {
	ID          uuid.UUID  `json:"id"`
	UserID      uuid.UUID  `json:"user_id"`
	AddressID   uuid.UUID  `json:"address_id"`
	AddressType int        `json:"address_type"`
	LastUsedOn  *time.Time `json:"last_used_on"`
	Address     *AddressDTO `json:"address,omitempty"`
}
