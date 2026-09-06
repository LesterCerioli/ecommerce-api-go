package pricing

import (
	"time"

	"github.com/google/uuid"
)

type CartRuleDTO struct {
	ID                    uuid.UUID  `json:"id"`
	Name                  string     `json:"name"`
	Description           string     `json:"description"`
	IsActive              bool       `json:"is_active"`
	StartOn               *time.Time `json:"start_on"`
	EndOn                 *time.Time `json:"end_on"`
	IsCouponRequired      bool       `json:"is_coupon_required"`
	RuleToApply           string     `json:"rule_to_apply"`
	DiscountAmount        float64    `json:"discount_amount"`
	MaxDiscountAmount     *float64   `json:"max_discount_amount"`
	DiscountStep          *int       `json:"discount_step"`
	UsageLimitPerCoupon   *int       `json:"usage_limit_per_coupon"`
	UsageLimitPerCustomer *int       `json:"usage_limit_per_customer"`
	CreatedOn             time.Time  `json:"created_on"`
	LatestUpdatedOn       time.Time  `json:"latest_updated_on"`
}
