package catalog

import (
	"time"

	"github.com/google/uuid"
)

type ProductPriceHistoryDTO struct {
	ID                uuid.UUID  `json:"id"`
	ProductID         uuid.UUID  `json:"product_id"`
	CreatedByID       uuid.UUID  `json:"created_by_id"`
	CreatedByName     string     `json:"created_by_name,omitempty"`
	CreatedOn         time.Time  `json:"created_on"`
	Price             *float64   `json:"price"`
	OldPrice          *float64   `json:"old_price"`
	SpecialPrice      *float64   `json:"special_price"`
	SpecialPriceStart *time.Time `json:"special_price_start"`
	SpecialPriceEnd   *time.Time `json:"special_price_end"`
}
