package shipping_table_rate

import "github.com/google/uuid"

type PriceAndDestinationDTO struct {
	ID                uuid.UUID  `json:"id"`
	CountryID         string     `json:"country_id"`
	CountryName       string     `json:"country_name,omitempty"`
	StateOrProvinceID *uuid.UUID `json:"state_or_province_id"`
	StateOrProvinceName string   `json:"state_or_province_name,omitempty"`
	DistrictID        *uuid.UUID `json:"district_id"`
	DistrictName      string     `json:"district_name,omitempty"`
	ZipCode           string     `json:"zip_code"`
	Note              string     `json:"note"`
	MinOrderSubtotal  float64    `json:"min_order_subtotal"`
	ShippingPrice     float64    `json:"shipping_price"`
}
