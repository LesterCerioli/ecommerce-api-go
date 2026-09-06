package core

import "github.com/google/uuid"

type StateOrProvinceDTO struct {
	ID        uuid.UUID `json:"id"`
	CountryID string    `json:"country_id"`
	CountryName string  `json:"country_name,omitempty"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
}
