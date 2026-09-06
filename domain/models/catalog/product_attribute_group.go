package catalog

import "github.com/google/uuid"

type ProductAttributeGroup struct {
	ID         uuid.UUID            `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name       string               `gorm:"column:name;size:450;not null"`
	Attributes []ProductAttribute   `gorm:"foreignKey:GroupID;references:ID"`
}

func (ProductAttributeGroup) TableName() string {
	return "public.product_attribute_groups"
}
