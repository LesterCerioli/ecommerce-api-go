package pricing

import (
	"time"

	"github.com/google/uuid"
)

type CatalogRule struct {
	ID                uuid.UUID                  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name              string                     `gorm:"column:name;size:450;not null"`
	Description       string                     `gorm:"column:description"`
	IsActive          bool                       `gorm:"column:is_active;not null"`
	StartOn           *time.Time                 `gorm:"column:start_on"`
	EndOn             *time.Time                 `gorm:"column:end_on"`
	RuleToApply       string                     `gorm:"column:rule_to_apply;size:450"`
	DiscountAmount    float64                    `gorm:"column:discount_amount;not null"`
	MaxDiscountAmount *float64                   `gorm:"column:max_discount_amount"`
	CustomerGroups    []CatalogRuleCustomerGroup  `gorm:"foreignKey:CatalogRuleID;references:ID"`
	CreatedOn         time.Time                  `gorm:"column:created_on;not null"`
	LatestUpdatedOn   time.Time                  `gorm:"column:latest_updated_on;not null"`
}

func (CatalogRule) TableName() string {
	return "public.catalog_rules"
}
