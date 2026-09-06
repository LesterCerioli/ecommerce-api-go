package tax

import "github.com/google/uuid"

type TaxClass struct {
	ID   uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name string    `gorm:"column:name;size:450;not null"`
}

func (TaxClass) TableName() string {
	return "public.tax_classes"
}
