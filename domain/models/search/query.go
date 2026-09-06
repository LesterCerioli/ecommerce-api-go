package search

import (
	"time"

	"github.com/google/uuid"
)

type Query struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	QueryText    string    `gorm:"column:query_text;size:500;not null"`
	ResultsCount int       `gorm:"column:results_count;not null"`
	CreatedOn    time.Time `gorm:"column:created_on;not null"`
}

func (Query) TableName() string {
	return "public.search_queries"
}
