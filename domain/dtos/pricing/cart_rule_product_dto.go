package pricing

import "github.com/google/uuid"

type CartRuleProductDTO struct {
	ProductID   uuid.UUID `json:"product_id"`
	ProductName string    `json:"product_name,omitempty"`
	CartRuleID  uuid.UUID `json:"cart_rule_id"`
}
