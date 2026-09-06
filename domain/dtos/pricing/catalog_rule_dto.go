package pricing

import (
	"time"

	"github.com/google/uuid"
)

type CatalogRuleDTO struct {
	ID                uuid.UUID  `json:"id"`
	Name              string     `json:"name"`
	Description       string     `json:"description"`
	IsActive          bool       `json:"is_active"`
	StartOn           *time.Time `json:"start_on"`
	EndOn             *time.Time `json:"end_on"`
	RuleToApply       string     `json:"rule_to_apply"`
	DiscountAmount    float64    `json:"discount_amount"`
	MaxDiscountAmount *float64   `json:"max_discount_amount"`
	CreatedOn         time.Time  `json:"created_on"`
	LatestUpdatedOn   time.Time  `json:"latest_updated_on"`
}
