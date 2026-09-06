package product_comparison

import (
	"time"

	"github.com/google/uuid"
)

type ComparingProductDTO struct {
	ID        uuid.UUID `json:"id"`
	CreatedOn time.Time `json:"created_on"`
	UserID    uuid.UUID `json:"user_id"`
	ProductID uuid.UUID `json:"product_id"`
}
