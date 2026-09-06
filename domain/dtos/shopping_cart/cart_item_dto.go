package shopping_cart

import (
	"time"

	"github.com/google/uuid"
)

type CartItemDTO struct {
	ID           uuid.UUID `json:"id"`
	CreatedOn    time.Time `json:"created_on"`
	ProductID    uuid.UUID `json:"product_id"`
	ProductName  string    `json:"product_name,omitempty"`
	ProductImage string    `json:"product_image,omitempty"`
	ProductPrice float64   `json:"product_price,omitempty"`
	Quantity     int       `json:"quantity"`
	CartID       uuid.UUID `json:"cart_id"`
}
