package search

import (
	"time"

	"github.com/google/uuid"
)

type QueryDTO struct {
	ID           uuid.UUID `json:"id"`
	QueryText    string    `json:"query_text"`
	ResultsCount int       `json:"results_count"`
	CreatedOn    time.Time `json:"created_on"`
}
