package shopping_cart

import (
	"time"

	"github.com/google/uuid"

	"ecommerce-api-go/domain/models/core"
)

type Cart struct {
	ID                     uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	CustomerID             uuid.UUID  `gorm:"column:customer_id;type:uuid;not null"`
	Customer               *core.User `gorm:"foreignKey:CustomerID;references:ID"`
	CreatedByID            uuid.UUID  `gorm:"column:created_by_id;type:uuid;not null"`
	CreatedBy              *core.User `gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedOn              time.Time  `gorm:"column:created_on;not null"`
	LatestUpdatedOn        time.Time  `gorm:"column:latest_updated_on;not null"`
	IsActive               bool       `gorm:"column:is_active;not null"`
	LockedOnCheckout       bool       `gorm:"column:locked_on_checkout;not null"`
	CouponCode             string     `gorm:"column:coupon_code;size:450"`
	CouponRuleName         string     `gorm:"column:coupon_rule_name;size:450"`
	ShippingMethod         string     `gorm:"column:shipping_method;size:450"`
	IsProductPriceIncludeTax bool     `gorm:"column:is_product_price_include_tax;not null"`
	ShippingAmount         *float64   `gorm:"column:shipping_amount"`
	TaxAmount              *float64   `gorm:"column:tax_amount"`
	Items                  []CartItem `gorm:"foreignKey:CartID;references:ID"`
	ShippingData           string     `gorm:"column:shipping_data"`
	OrderNote              string     `gorm:"column:order_note;size:1000"`
}

func (Cart) TableName() string {
	return "public.carts"
}
