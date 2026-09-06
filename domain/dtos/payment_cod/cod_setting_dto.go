package payment_cod

type CoDSettingDTO struct {
	MinOrderValue *float64 `json:"min_order_value"`
	MaxOrderValue *float64 `json:"max_order_value"`
	PaymentFee    float64  `json:"payment_fee"`
}
