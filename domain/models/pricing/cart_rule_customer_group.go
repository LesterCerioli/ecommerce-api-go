package pricing

import "github.com/google/uuid"

type CartRuleCustomerGroup struct {
	CartRuleID      uuid.UUID `gorm:"primaryKey;column:cart_rule_id;type:uuid"`
	CartRule        *CartRule `gorm:"foreignKey:CartRuleID;references:ID"`
	CustomerGroupID uuid.UUID `gorm:"primaryKey;column:customer_group_id;type:uuid"`
}

func (CartRuleCustomerGroup) TableName() string {
	return "public.cart_rule_customer_groups"
}
