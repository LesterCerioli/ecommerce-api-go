package orders

import "github.com/google/uuid"

type OrderItem struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	OrderID      uuid.UUID `gorm:"column:order_id;type:uuid;not null"`
	Order        *Order    `gorm:"foreignKey:OrderID;references:ID"`
	ProductID    uuid.UUID `gorm:"column:product_id;type:uuid;not null"`
	ProductPrice float64   `gorm:"column:product_price;not null"`
	Quantity     int       `gorm:"column:quantity;not null"`
	DiscountAmount float64 `gorm:"column:discount_amount;not null"`
	TaxAmount    float64   `gorm:"column:tax_amount;not null"`
	TaxPercent   float64   `gorm:"column:tax_percent;not null"`
}

func (OrderItem) TableName() string {
	return "public.order_items"
}
