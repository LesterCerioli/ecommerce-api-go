package payment_cod

type CoDSetting struct {
	MinOrderValue *float64 `gorm:"column:min_order_value"`
	MaxOrderValue *float64 `gorm:"column:max_order_value"`
	PaymentFee    float64  `gorm:"column:payment_fee;not null"`
}

func (CoDSetting) TableName() string {
	return "public.cod_settings"
}
