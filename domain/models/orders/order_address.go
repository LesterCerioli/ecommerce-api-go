package orders

import "github.com/google/uuid"

type OrderAddress struct {
	ID                uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	ContactName       string    `gorm:"column:contact_name;size:450"`
	Phone             string    `gorm:"column:phone;size:450"`
	AddressLine1      string    `gorm:"column:address_line_1;size:450"`
	AddressLine2      string    `gorm:"column:address_line_2;size:450"`
	City              string    `gorm:"column:city;size:450"`
	ZipCode           string    `gorm:"column:zip_code;size:450"`
	DistrictID        *uuid.UUID `gorm:"column:district_id;type:uuid"`
	StateOrProvinceID uuid.UUID `gorm:"column:state_or_province_id;type:uuid;not null"`
	CountryID         string    `gorm:"column:country_id;size:450;not null"`
}

func (OrderAddress) TableName() string {
	return "public.order_addresses"
}
