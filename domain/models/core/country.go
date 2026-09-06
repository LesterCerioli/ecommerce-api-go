package core

type Country struct {
	ID                string           `gorm:"primaryKey;column:id;size:450"`
	Name              string           `gorm:"column:name;not null"`
	Code3             string           `gorm:"column:code3"`
	IsBillingEnabled  bool             `gorm:"column:is_billing_enabled;not null"`
	IsShippingEnabled bool             `gorm:"column:is_shipping_enabled;not null"`
	IsCityEnabled     bool             `gorm:"column:is_city_enabled;not null"`
	IsZipCodeEnabled  bool             `gorm:"column:is_zip_code_enabled;not null"`
	IsDistrictEnabled bool             `gorm:"column:is_district_enabled;not null"`
	StatesOrProvinces []StateOrProvince `gorm:"foreignKey:CountryID;references:ID"`
}

func (Country) TableName() string {
	return "public.countries"
}
