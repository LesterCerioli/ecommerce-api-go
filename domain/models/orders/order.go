package orders

import (
	"time"

	"github.com/google/uuid"

	sharedorders "ecommerce-api-go/domain/shared/domain/orders"
	"ecommerce-api-go/domain/models/core"
)

type OrderStatus = sharedorders.OrderStatus

type Order struct {
	ID                uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	CustomerID        uuid.UUID      `gorm:"column:customer_id;type:uuid;not null"`
	Customer          *core.User     `gorm:"foreignKey:CustomerID;references:ID"`
	LatestUpdatedOn   time.Time      `gorm:"column:latest_updated_on;not null"`
	LatestUpdatedByID uuid.UUID      `gorm:"column:latest_updated_by_id;type:uuid;not null"`
	LatestUpdatedBy   *core.User     `gorm:"foreignKey:LatestUpdatedByID;references:ID"`
	CreatedOn         time.Time      `gorm:"column:created_on;not null"`
	CreatedByID       uuid.UUID      `gorm:"column:created_by_id;type:uuid;not null"`
	CreatedBy         *core.User     `gorm:"foreignKey:CreatedByID;references:ID"`
	VendorID          *uuid.UUID     `gorm:"column:vendor_id;type:uuid"`
	CouponCode        string         `gorm:"column:coupon_code;size:450"`
	CouponRuleName    string         `gorm:"column:coupon_rule_name;size:450"`
	DiscountAmount    float64        `gorm:"column:discount_amount;not null"`
	SubTotal          float64        `gorm:"column:sub_total;not null"`
	SubTotalWithDiscount float64     `gorm:"column:sub_total_with_discount;not null"`
	ShippingAddressID uuid.UUID     `gorm:"column:shipping_address_id;type:uuid;not null"`
	ShippingAddress   *OrderAddress  `gorm:"foreignKey:ShippingAddressID;references:ID"`
	BillingAddressID  uuid.UUID     `gorm:"column:billing_address_id;type:uuid;not null"`
	BillingAddress    *OrderAddress  `gorm:"foreignKey:BillingAddressID;references:ID"`
	OrderItems        []OrderItem    `gorm:"foreignKey:OrderID;references:ID"`
	OrderStatus       OrderStatus    `gorm:"column:order_status;not null"`
	OrderNote         string         `gorm:"column:order_note;size:1000"`
	ParentID          *uuid.UUID     `gorm:"column:parent_id;type:uuid"`
	Parent            *Order         `gorm:"foreignKey:ParentID;references:ID"`
	IsMasterOrder     bool           `gorm:"column:is_master_order;not null"`
	ShippingMethod    string         `gorm:"column:shipping_method;size:450"`
	ShippingFeeAmount float64        `gorm:"column:shipping_fee_amount;not null"`
	TaxAmount         float64        `gorm:"column:tax_amount;not null"`
	OrderTotal        float64        `gorm:"column:order_total;not null"`
	PaymentMethod     string         `gorm:"column:payment_method;size:450"`
	PaymentFeeAmount  float64        `gorm:"column:payment_fee_amount;not null"`
	Children          []Order        `gorm:"foreignKey:ParentID;references:ID"`
}

func (Order) TableName() string {
	return "public.orders"
}
