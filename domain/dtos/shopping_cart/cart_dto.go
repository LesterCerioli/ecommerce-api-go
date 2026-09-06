package shopping_cart

import (
	"time"

	"github.com/google/uuid"
)

type CartDTO struct {
	ID                       uuid.UUID  `json:"id"`
	CustomerID               uuid.UUID  `json:"customer_id"`
	CustomerName             string     `json:"customer_name,omitempty"`
	CreatedByID              uuid.UUID  `json:"created_by_id"`
	CreatedOn                time.Time  `json:"created_on"`
	LatestUpdatedOn          time.Time  `json:"latest_updated_on"`
	IsActive                 bool       `json:"is_active"`
	LockedOnCheckout         bool       `json:"locked_on_checkout"`
	CouponCode               string     `json:"coupon_code"`
	CouponRuleName           string     `json:"coupon_rule_name"`
	ShippingMethod           string     `json:"shipping_method"`
	IsProductPriceIncludeTax bool       `json:"is_product_price_include_tax"`
	ShippingAmount           *float64   `json:"shipping_amount"`
	TaxAmount                *float64   `json:"tax_amount"`
	ShippingData             string     `json:"shipping_data"`
	OrderNote                string     `json:"order_note"`
	Items                    []CartItemDTO `json:"items,omitempty"`
}
