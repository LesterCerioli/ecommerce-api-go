package shipping_free

type FreeShippingSetting struct {
	MinimumOrderAmount float64 `gorm:"column:minimum_order_amount;not null"`
}

func (FreeShippingSetting) TableName() string {
	return "public.free_shipping_settings"
}
