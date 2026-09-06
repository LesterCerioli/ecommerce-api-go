package catalog

import (
	"github.com/google/uuid"

	"ecommerce-api-go/domain/models/core"
)

type ProductMedia struct {
	ID           uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	ProductID    uuid.UUID  `gorm:"column:product_id;type:uuid;not null"`
	Product      *Product   `gorm:"foreignKey:ProductID;references:ID"`
	MediaID      uuid.UUID  `gorm:"column:media_id;type:uuid;not null"`
	Media        *core.Media `gorm:"foreignKey:MediaID;references:ID"`
	DisplayOrder int        `gorm:"column:display_order;not null"`
}

func (ProductMedia) TableName() string {
	return "public.product_media"
}
