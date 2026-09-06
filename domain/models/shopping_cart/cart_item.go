package shopping_cart

import (
	"time"

	"github.com/google/uuid"

	"ecommerce-api-go/domain/models/catalog"
)

type CartItem struct {
	ID        time.Time      `gorm:"primaryKey;column:created_on"`
	CreatedOn time.Time      `gorm:"column:created_on;not null"`
	ProductID uuid.UUID      `gorm:"column:product_id;type:uuid;not null"`
	Product   *catalog.Product `gorm:"foreignKey:ProductID;references:ID"`
	Quantity  int            `gorm:"column:quantity;not null"`
	CartID    uuid.UUID      `gorm:"column:cart_id;type:uuid;not null"`
	Cart      *Cart          `gorm:"foreignKey:CartID;references:ID"`
}

func (CartItem) TableName() string {
	return "public.cart_items"
}
