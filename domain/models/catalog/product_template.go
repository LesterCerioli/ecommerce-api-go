package catalog

import "github.com/google/uuid"

type ProductTemplate struct {
	ID               uuid.UUID                       `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name             string                          `gorm:"column:name;size:450;not null"`
	ProductAttributes []ProductTemplateProductAttribute `gorm:"foreignKey:ProductTemplateID;references:ID"`
}

func (ProductTemplate) TableName() string {
	return "public.product_templates"
}
