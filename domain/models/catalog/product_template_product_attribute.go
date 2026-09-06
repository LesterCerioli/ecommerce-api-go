package catalog

import "github.com/google/uuid"

type ProductTemplateProductAttribute struct {
	ProductTemplateID  uuid.UUID          `gorm:"primaryKey;column:product_template_id;type:uuid"`
	ProductTemplate    *ProductTemplate   `gorm:"foreignKey:ProductTemplateID;references:ID"`
	ProductAttributeID uuid.UUID          `gorm:"primaryKey;column:product_attribute_id;type:uuid"`
	ProductAttribute   *ProductAttribute  `gorm:"foreignKey:ProductAttributeID;references:ID"`
}

func (ProductTemplateProductAttribute) TableName() string {
	return "public.product_template_product_attributes"
}
