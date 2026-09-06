package catalog

import "github.com/google/uuid"

type ProductCategory struct {
	ID                uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	IsFeaturedProduct bool       `gorm:"column:is_featured_product;not null"`
	DisplayOrder      int        `gorm:"column:display_order;not null"`
	CategoryID        uuid.UUID  `gorm:"column:category_id;type:uuid;not null"`
	Category          *Category  `gorm:"foreignKey:CategoryID;references:ID"`
	ProductID         uuid.UUID  `gorm:"column:product_id;type:uuid;not null"`
	Product           *Product   `gorm:"foreignKey:ProductID;references:ID"`
}

func (ProductCategory) TableName() string {
	return "public.product_categories"
}
