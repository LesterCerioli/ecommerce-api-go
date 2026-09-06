package payments

type PaymentProvider struct {
	ID                        string `gorm:"primaryKey;column:id;size:450"`
	Name                      string `gorm:"column:name;size:450;not null"`
	IsEnabled                 bool   `gorm:"column:is_enabled;not null"`
	ConfigureURL              string `gorm:"column:configure_url;size:450"`
	LandingViewComponentName  string `gorm:"column:landing_view_component_name;size:450"`
	AdditionalSettings        string `gorm:"column:additional_settings"`
}

func (PaymentProvider) TableName() string {
	return "public.payment_providers"
}
