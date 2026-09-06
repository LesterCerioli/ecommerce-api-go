package catalog

import "github.com/google/uuid"

type ProductOption struct {
	ID   uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name string    `gorm:"column:name;size:450;not null"`
}

func (ProductOption) TableName() string {
	return "public.product_options"
}
