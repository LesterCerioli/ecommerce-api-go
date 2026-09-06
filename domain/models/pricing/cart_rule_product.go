package pricing

import "github.com/google/uuid"

type CartRuleProduct struct {
	ProductID  uuid.UUID `gorm:"primaryKey;column:product_id;type:uuid"`
	CartRuleID uuid.UUID `gorm:"primaryKey;column:cart_rule_id;type:uuid"`
	CartRule   *CartRule `gorm:"foreignKey:CartRuleID;references:ID"`
}

func (CartRuleProduct) TableName() string {
	return "public.cart_rule_products"
}
