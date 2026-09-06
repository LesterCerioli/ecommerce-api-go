package pricing

import (
	"time"

	"github.com/google/uuid"
)

type CartRule struct {
	ID                    uuid.UUID                  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name                  string                     `gorm:"column:name;size:450;not null"`
	Description           string                     `gorm:"column:description"`
	IsActive              bool                       `gorm:"column:is_active;not null"`
	StartOn               *time.Time                 `gorm:"column:start_on"`
	EndOn                 *time.Time                 `gorm:"column:end_on"`
	IsCouponRequired      bool                       `gorm:"column:is_coupon_required;not null"`
	RuleToApply           string                     `gorm:"column:rule_to_apply;size:450"`
	DiscountAmount        float64                    `gorm:"column:discount_amount;not null"`
	MaxDiscountAmount     *float64                   `gorm:"column:max_discount_amount"`
	DiscountStep          *int                       `gorm:"column:discount_step"`
	UsageLimitPerCoupon   *int                       `gorm:"column:usage_limit_per_coupon"`
	UsageLimitPerCustomer *int                       `gorm:"column:usage_limit_per_customer"`
	Coupons               []Coupon                   `gorm:"foreignKey:CartRuleID;references:ID"`
	CustomerGroups        []CartRuleCustomerGroup     `gorm:"foreignKey:CartRuleID;references:ID"`
	Products              []CartRuleProduct           `gorm:"foreignKey:CartRuleID;references:ID"`
	Categories            []CartRuleCategory          `gorm:"foreignKey:CartRuleID;references:ID"`
	CreatedOn             time.Time                  `gorm:"column:created_on;not null"`
	LatestUpdatedOn       time.Time                  `gorm:"column:latest_updated_on;not null"`
}

func (CartRule) TableName() string {
	return "public.cart_rules"
}
