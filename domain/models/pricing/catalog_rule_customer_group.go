package pricing

import "github.com/google/uuid"

type CatalogRuleCustomerGroup struct {
	CatalogRuleID   uuid.UUID `gorm:"primaryKey;column:catalog_rule_id;type:uuid"`
	CatalogRule     *CatalogRule `gorm:"foreignKey:CatalogRuleID;references:ID"`
	CustomerGroupID uuid.UUID `gorm:"primaryKey;column:customer_group_id;type:uuid"`
}

func (CatalogRuleCustomerGroup) TableName() string {
	return "public.catalog_rule_customer_groups"
}
