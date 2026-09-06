package orders

import (
	"time"

	"github.com/google/uuid"
)

type OrderDTO struct {
	ID                    uuid.UUID  `json:"id"`
	CustomerID            uuid.UUID  `json:"customer_id"`
	CustomerName          string     `json:"customer_name,omitempty"`
	LatestUpdatedOn       time.Time  `json:"latest_updated_on"`
	LatestUpdatedByID     uuid.UUID  `json:"latest_updated_by_id"`
	CreatedOn             time.Time  `json:"created_on"`
	CreatedByID           uuid.UUID  `json:"created_by_id"`
	VendorID              *uuid.UUID `json:"vendor_id"`
	CouponCode            string     `json:"coupon_code"`
	CouponRuleName        string     `json:"coupon_rule_name"`
	DiscountAmount        float64    `json:"discount_amount"`
	SubTotal              float64    `json:"sub_total"`
	SubTotalWithDiscount  float64    `json:"sub_total_with_discount"`
	ShippingAddressID     uuid.UUID  `json:"shipping_address_id"`
	BillingAddressID      uuid.UUID  `json:"billing_address_id"`
	OrderStatus           int        `json:"order_status"`
	OrderNote             string     `json:"order_note"`
	ParentID              *uuid.UUID `json:"parent_id"`
	IsMasterOrder         bool       `json:"is_master_order"`
	ShippingMethod        string     `json:"shipping_method"`
	ShippingFeeAmount     float64    `json:"shipping_fee_amount"`
	TaxAmount             float64    `json:"tax_amount"`
	OrderTotal            float64    `json:"order_total"`
	PaymentMethod         string     `json:"payment_method"`
	PaymentFeeAmount      float64    `json:"payment_fee_amount"`
	OrderItems            []OrderItemDTO `json:"order_items,omitempty"`
	ShippingAddress       *OrderAddressDTO `json:"shipping_address,omitempty"`
	BillingAddress        *OrderAddressDTO `json:"billing_address,omitempty"`
}
