package orders

import "github.com/google/uuid"

type OrderItemDTO struct {
	ID             uuid.UUID `json:"id"`
	OrderID        uuid.UUID `json:"order_id"`
	ProductID      uuid.UUID `json:"product_id"`
	ProductName    string    `json:"product_name,omitempty"`
	ProductImage   string    `json:"product_image,omitempty"`
	ProductPrice   float64   `json:"product_price"`
	Quantity       int       `json:"quantity"`
	DiscountAmount float64   `json:"discount_amount"`
	TaxAmount      float64   `json:"tax_amount"`
	TaxPercent     float64   `json:"tax_percent"`
}
