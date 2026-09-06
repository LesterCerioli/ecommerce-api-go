package pricing

import (
	"time"

	"github.com/google/uuid"
)

type Coupon struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	CartRuleID  uuid.UUID `gorm:"column:cart_rule_id;type:uuid;not null"`
	CartRule    *CartRule `gorm:"foreignKey:CartRuleID;references:ID"`
	Code        string    `gorm:"column:code;size:450;not null"`
	CreatedOn   time.Time `gorm:"column:created_on;not null"`
}

func (Coupon) TableName() string {
	return "public.coupons"
}
