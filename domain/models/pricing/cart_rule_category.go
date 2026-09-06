package pricing

import "github.com/google/uuid"

type CartRuleCategory struct {
	CategoryID uuid.UUID `gorm:"primaryKey;column:category_id;type:uuid"`
	CartRuleID uuid.UUID `gorm:"primaryKey;column:cart_rule_id;type:uuid"`
	CartRule   *CartRule `gorm:"foreignKey:CartRuleID;references:ID"`
}

func (CartRuleCategory) TableName() string {
	return "public.cart_rule_categories"
}
