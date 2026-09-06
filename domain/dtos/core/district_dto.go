package core

import "github.com/google/uuid"

type DistrictDTO struct {
	ID                uuid.UUID `json:"id"`
	StateOrProvinceID uuid.UUID `json:"state_or_province_id"`
	StateOrProvinceName string  `json:"state_or_province_name,omitempty"`
	Name              string    `json:"name"`
	Type              string    `json:"type"`
	Location          string    `json:"location"`
}
