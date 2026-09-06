package shipping

type ShippingProviderDTO struct {
	ID                                    string `json:"id"`
	Name                                  string `json:"name"`
	IsEnabled                             bool   `json:"is_enabled"`
	ConfigureURL                          string `json:"configure_url"`
	ToAllShippingEnabledCountries         bool   `json:"to_all_shipping_enabled_countries"`
	OnlyCountryIDsString                  string `json:"only_country_ids_string"`
	ToAllShippingEnabledStatesOrProvinces bool   `json:"to_all_shipping_enabled_states_or_provinces"`
	OnlyStateOrProvinceIDsString          string `json:"only_state_or_province_ids_string"`
	AdditionalSettings                    string `json:"additional_settings"`
	ShippingPriceServiceTypeName          string `json:"shipping_price_service_type_name"`
}
