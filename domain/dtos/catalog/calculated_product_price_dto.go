package catalog

import "github.com/google/uuid"

type CalculatedProductPriceDTO struct {
	ProductID      uuid.UUID `json:"product_id"`
	Price          float64   `json:"price"`
	OldPrice       *float64  `json:"old_price"`
	PercentOfSaving float64  `json:"percent_of_saving"`
	PriceString    string    `json:"price_string"`
	OldPriceString string    `json:"old_price_string"`
}
