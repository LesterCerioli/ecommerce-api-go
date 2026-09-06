package pricing

import (
	"time"

	"github.com/google/uuid"
)

type CartRuleUsageDTO struct {
	ID          uuid.UUID  `json:"id"`
	CartRuleID  uuid.UUID  `json:"cart_rule_id"`
	CouponID    *uuid.UUID `json:"coupon_id"`
	UserID      uuid.UUID  `json:"user_id"`
	UserName    string     `json:"user_name,omitempty"`
	OrderID     uuid.UUID  `json:"order_id"`
	CreatedOn   time.Time  `json:"created_on"`
}
