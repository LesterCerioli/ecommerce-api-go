package catalog

import "github.com/google/uuid"

type ProductOptionCombination struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	ProductID uuid.UUID      `gorm:"column:product_id;type:uuid;not null"`
	Product   *Product       `gorm:"foreignKey:ProductID;references:ID"`
	OptionID  uuid.UUID      `gorm:"column:option_id;type:uuid;not null"`
	Option    *ProductOption `gorm:"foreignKey:OptionID;references:ID"`
	Value     string         `gorm:"column:value;size:450"`
	SortIndex int            `gorm:"column:sort_index;not null"`
}

func (ProductOptionCombination) TableName() string {
	return "public.product_option_combinations"
}
