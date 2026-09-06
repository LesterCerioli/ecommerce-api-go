package catalog

import (
	"time"

	"github.com/google/uuid"

	"ecommerce-api-go/domain/models/core"
)

type ProductPriceHistory struct {
	ID               uuid.UUID   `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	ProductID        uuid.UUID   `gorm:"column:product_id;type:uuid;not null"`
	Product          *Product    `gorm:"foreignKey:ProductID;references:ID"`
	CreatedByID      uuid.UUID   `gorm:"column:created_by_id;type:uuid;not null"`
	CreatedBy        *core.User  `gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedOn        time.Time   `gorm:"column:created_on;not null"`
	Price            *float64    `gorm:"column:price"`
	OldPrice         *float64    `gorm:"column:old_price"`
	SpecialPrice     *float64    `gorm:"column:special_price"`
	SpecialPriceStart *time.Time `gorm:"column:special_price_start"`
	SpecialPriceEnd  *time.Time  `gorm:"column:special_price_end"`
}

func (ProductPriceHistory) TableName() string {
	return "public.product_price_histories"
}
