package product_comparison

import (
	"time"

	"github.com/google/uuid"

	"ecommerce-api-go/domain/models/core"
)

type ComparingProduct struct {
	ID        uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	CreatedOn time.Time  `gorm:"column:created_on;not null"`
	UserID    uuid.UUID  `gorm:"column:user_id;type:uuid;not null"`
	User      *core.User `gorm:"foreignKey:UserID;references:ID"`
	ProductID uuid.UUID  `gorm:"column:product_id;type:uuid;not null"`
}

func (ComparingProduct) TableName() string {
	return "public.comparing_products"
}
