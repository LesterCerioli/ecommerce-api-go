package wishlist

import (
	"time"

	"github.com/google/uuid"
)

type WishListDTO struct {
	ID              uuid.UUID      `json:"id"`
	UserID          uuid.UUID      `json:"user_id"`
	UserName        string         `json:"user_name,omitempty"`
	SharingCode     string         `json:"sharing_code"`
	Items           []WishListItemDTO `json:"items,omitempty"`
	CreatedOn       time.Time      `json:"created_on"`
	LatestUpdatedOn time.Time      `json:"latest_updated_on"`
}
