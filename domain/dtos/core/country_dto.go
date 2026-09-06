package core

type CountryDTO struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Code3             string `json:"code3"`
	IsBillingEnabled  bool   `json:"is_billing_enabled"`
	IsShippingEnabled bool   `json:"is_shipping_enabled"`
	IsCityEnabled     bool   `json:"is_city_enabled"`
	IsZipCodeEnabled  bool   `json:"is_zip_code_enabled"`
	IsDistrictEnabled bool   `json:"is_district_enabled"`
}
