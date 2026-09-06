package core

import "github.com/google/uuid"

type Address struct {
	ID                uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	ContactName       string     `gorm:"column:contact_name;not null"`
	Phone             string     `gorm:"column:phone;not null"`
	AddressLine1      string     `gorm:"column:address_line_1;not null"`
	AddressLine2      string     `gorm:"column:address_line_2"`
	City              string     `gorm:"column:city;not null"`
	ZipCode           string     `gorm:"column:zip_code;not null"`
	DistrictID        *uuid.UUID `gorm:"column:district_id;type:uuid"`
	District          *District  `gorm:"foreignKey:DistrictID;references:ID"`
	StateOrProvinceID uuid.UUID  `gorm:"column:state_or_province_id;type:uuid;not null"`
	StateOrProvince   *StateOrProvince `gorm:"foreignKey:StateOrProvinceID;references:ID"`
	CountryID         string     `gorm:"column:country_id;not null"`
	Country           *Country   `gorm:"foreignKey:CountryID;references:ID"`
	UserAddresses     []UserAddress `gorm:"foreignKey:AddressID;references:ID"`
}

func (Address) TableName() string {
	return "public.addresses"
}
