package wishlist

import (
	"time"

	"github.com/google/uuid"
)

type WishListItemDTO struct {
	ID              uuid.UUID `json:"id"`
	WishListID      uuid.UUID `json:"wish_list_id"`
	ProductID       uuid.UUID `json:"product_id"`
	ProductName     string    `json:"product_name,omitempty"`
	ProductImage    string    `json:"product_image,omitempty"`
	Description     string    `json:"description"`
	Quantity        int       `json:"quantity"`
	CreatedOn       time.Time `json:"created_on"`
	LatestUpdatedOn time.Time `json:"latest_updated_on"`
}
