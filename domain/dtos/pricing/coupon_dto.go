package pricing

import (
	"time"

	"github.com/google/uuid"
)

type CouponDTO struct {
	ID          uuid.UUID `json:"id"`
	CartRuleID  uuid.UUID `json:"cart_rule_id"`
	Code        string    `json:"code"`
	CreatedOn   time.Time `json:"created_on"`
}
