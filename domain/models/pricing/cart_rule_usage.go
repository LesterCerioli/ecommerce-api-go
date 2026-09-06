package pricing

import (
	"time"

	"github.com/google/uuid"

	"ecommerce-api-go/domain/models/core"
)

type CartRuleUsage struct {
	ID          uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	CartRuleID  uuid.UUID  `gorm:"column:cart_rule_id;type:uuid;not null"`
	CartRule    *CartRule  `gorm:"foreignKey:CartRuleID;references:ID"`
	CouponID    *uuid.UUID `gorm:"column:coupon_id;type:uuid"`
	Coupon      *Coupon    `gorm:"foreignKey:CouponID;references:ID"`
	UserID      uuid.UUID  `gorm:"column:user_id;type:uuid;not null"`
	User        *core.User `gorm:"foreignKey:UserID;references:ID"`
	OrderID     uuid.UUID  `gorm:"column:order_id;type:uuid;not null"`
	CreatedOn   time.Time  `gorm:"column:created_on;not null"`
}

func (CartRuleUsage) TableName() string {
	return "public.cart_rule_usages"
}
