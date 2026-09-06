package orders

import "github.com/google/uuid"

type OrderAddressDTO struct {
	ID                uuid.UUID  `json:"id"`
	ContactName       string     `json:"contact_name"`
	Phone             string     `json:"phone"`
	AddressLine1      string     `json:"address_line_1"`
	AddressLine2      string     `json:"address_line_2"`
	City              string     `json:"city"`
	ZipCode           string     `json:"zip_code"`
	DistrictID        *uuid.UUID `json:"district_id"`
	StateOrProvinceID uuid.UUID  `json:"state_or_province_id"`
	CountryID         string     `json:"country_id"`
}
