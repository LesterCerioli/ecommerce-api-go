package catalog

import "github.com/google/uuid"

type ProductAttribute struct {
	ID             uuid.UUID                   `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name           string                      `gorm:"column:name;size:450;not null"`
	GroupID        uuid.UUID                   `gorm:"column:group_id;type:uuid;not null"`
	Group          *ProductAttributeGroup      `gorm:"foreignKey:GroupID;references:ID"`
	ProductTemplates []ProductTemplateProductAttribute `gorm:"foreignKey:ProductAttributeID;references:ID"`
}

func (ProductAttribute) TableName() string {
	return "public.product_attributes"
}
