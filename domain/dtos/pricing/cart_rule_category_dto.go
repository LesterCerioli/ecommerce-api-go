package pricing

import "github.com/google/uuid"

type CartRuleCategoryDTO struct {
	CategoryID   uuid.UUID `json:"category_id"`
	CategoryName string    `json:"category_name,omitempty"`
	CartRuleID   uuid.UUID `json:"cart_rule_id"`
}
