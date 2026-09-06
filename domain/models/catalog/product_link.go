package catalog

import (
	"github.com/google/uuid"

	sharedcatalog "ecommerce-api-go/domain/shared/domain/catalog"
)

type ProductLinkType = sharedcatalog.ProductLinkType

type ProductLink struct {
	ID             uuid.UUID       `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	ProductID      uuid.UUID       `gorm:"column:product_id;type:uuid;not null"`
	Product        *Product        `gorm:"foreignKey:ProductID;references:ID"`
	LinkedProductID uuid.UUID      `gorm:"column:linked_product_id;type:uuid;not null"`
	LinkedProduct  *Product        `gorm:"foreignKey:LinkedProductID;references:ID"`
	LinkType       ProductLinkType `gorm:"column:link_type;not null"`
}

func (ProductLink) TableName() string {
	return "public.product_links"
}
