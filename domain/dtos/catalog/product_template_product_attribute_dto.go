package catalog

import "github.com/google/uuid"

type ProductTemplateProductAttributeDTO struct {
	ProductTemplateID  uuid.UUID `json:"product_template_id"`
	ProductTemplateName string  `json:"product_template_name,omitempty"`
	ProductAttributeID uuid.UUID `json:"product_attribute_id"`
	ProductAttributeName string `json:"product_attribute_name,omitempty"`
}
