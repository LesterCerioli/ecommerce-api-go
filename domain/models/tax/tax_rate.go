package tax

import "github.com/google/uuid"

type TaxRate struct {
	ID                uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	TaxClassID        uuid.UUID  `gorm:"column:tax_class_id;type:uuid;not null"`
	TaxClass          *TaxClass  `gorm:"foreignKey:TaxClassID;references:ID"`
	CountryID         string     `gorm:"column:country_id;size:450"`
	StateOrProvinceID *uuid.UUID `gorm:"column:state_or_province_id;type:uuid"`
	Rate              float64    `gorm:"column:rate;not null"`
	ZipCode           string     `gorm:"column:zip_code;size:450"`
}

func (TaxRate) TableName() string {
	return "public.tax_rates"
}
