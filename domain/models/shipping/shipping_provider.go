package shipping

type ShippingProvider struct {
	ID                               string `gorm:"primaryKey;column:id;size:450"`
	Name                             string `gorm:"column:name;size:450;not null"`
	IsEnabled                        bool   `gorm:"column:is_enabled;not null"`
	ConfigureURL                     string `gorm:"column:configure_url;size:450"`
	ToAllShippingEnabledCountries    bool   `gorm:"column:to_all_shipping_enabled_countries;not null"`
	OnlyCountryIDsString             string `gorm:"column:only_country_ids_string;size:1000"`
	ToAllShippingEnabledStatesOrProvinces bool `gorm:"column:to_all_shipping_enabled_states_or_provinces;not null"`
	OnlyStateOrProvinceIDsString     string `gorm:"column:only_state_or_province_ids_string;size:1000"`
	AdditionalSettings               string `gorm:"column:additional_settings"`
	ShippingPriceServiceTypeName     string `gorm:"column:shipping_price_service_type_name;size:450"`
}

func (ShippingProvider) TableName() string {
	return "public.shipping_providers"
}
