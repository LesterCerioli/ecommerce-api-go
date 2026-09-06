package catalog

import "github.com/google/uuid"

type ProductOptionValue struct {
	ID          uuid.UUID       `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	OptionID    uuid.UUID       `gorm:"column:option_id;type:uuid;not null"`
	Option      *ProductOption  `gorm:"foreignKey:OptionID;references:ID"`
	ProductID   uuid.UUID       `gorm:"column:product_id;type:uuid;not null"`
	Product     *Product        `gorm:"foreignKey:ProductID;references:ID"`
	Value       string          `gorm:"column:value;size:450"`
	DisplayType string          `gorm:"column:display_type;size:450"`
	SortIndex   int             `gorm:"column:sort_index;not null"`
}

func (ProductOptionValue) TableName() string {
	return "public.product_option_values"
}
