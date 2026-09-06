package catalog

import "github.com/google/uuid"

type ProductAttributeValue struct {
	ID          uuid.UUID         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	AttributeID uuid.UUID         `gorm:"column:attribute_id;type:uuid;not null"`
	Attribute   *ProductAttribute `gorm:"foreignKey:AttributeID;references:ID"`
	ProductID   uuid.UUID         `gorm:"column:product_id;type:uuid;not null"`
	Product     *Product          `gorm:"foreignKey:ProductID;references:ID"`
	Value       string            `gorm:"column:value"`
}

func (ProductAttributeValue) TableName() string {
	return "public.product_attribute_values"
}
