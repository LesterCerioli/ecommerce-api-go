package payments

type PaymentProviderDTO struct {
	ID                       string `json:"id"`
	Name                     string `json:"name"`
	IsEnabled                bool   `json:"is_enabled"`
	ConfigureURL             string `json:"configure_url"`
	LandingViewComponentName string `json:"landing_view_component_name"`
	AdditionalSettings       string `json:"additional_settings"`
}
