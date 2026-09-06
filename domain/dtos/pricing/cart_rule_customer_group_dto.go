package pricing

import "github.com/google/uuid"

type CartRuleCustomerGroupDTO struct {
	CartRuleID      uuid.UUID `json:"cart_rule_id"`
	CustomerGroupID uuid.UUID `json:"customer_group_id"`
	CustomerGroupName string  `json:"customer_group_name,omitempty"`
}
