package shipping_table_rate

import "github.com/google/uuid"

type PriceAndDestination struct {
	ID                uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	CountryID         string     `gorm:"column:country_id;size:450"`
	StateOrProvinceID *uuid.UUID `gorm:"column:state_or_province_id;type:uuid"`
	DistrictID        *uuid.UUID `gorm:"column:district_id;type:uuid"`
	ZipCode           string     `gorm:"column:zip_code;size:450"`
	Note              string     `gorm:"column:note"`
	MinOrderSubtotal  float64    `gorm:"column:min_order_subtotal;not null"`
	ShippingPrice     float64    `gorm:"column:shipping_price;not null"`
}

func (PriceAndDestination) TableName() string {
	return "public.price_and_destinations"
}
