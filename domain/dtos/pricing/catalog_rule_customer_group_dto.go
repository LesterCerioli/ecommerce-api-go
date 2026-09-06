package pricing

import "github.com/google/uuid"

type CatalogRuleCustomerGroupDTO struct {
	CatalogRuleID   uuid.UUID `json:"catalog_rule_id"`
	CustomerGroupID uuid.UUID `json:"customer_group_id"`
	CustomerGroupName string  `json:"customer_group_name,omitempty"`
}
