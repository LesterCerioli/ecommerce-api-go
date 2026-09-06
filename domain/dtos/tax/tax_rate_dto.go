package tax

import "github.com/google/uuid"

type TaxRateDTO struct {
	ID                uuid.UUID  `json:"id"`
	TaxClassID        uuid.UUID  `json:"tax_class_id"`
	TaxClassName      string     `json:"tax_class_name,omitempty"`
	CountryID         string     `json:"country_id"`
	CountryName       string     `json:"country_name,omitempty"`
	StateOrProvinceID *uuid.UUID `json:"state_or_province_id"`
	StateOrProvinceName string   `json:"state_or_province_name,omitempty"`
	Rate              float64    `json:"rate"`
	ZipCode           string     `json:"zip_code"`
}
