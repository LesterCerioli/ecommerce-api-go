package inventory

import "github.com/google/uuid"

type Warehouse struct {
	ID        uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name      string     `gorm:"column:name;size:450;not null"`
	AddressID uuid.UUID  `gorm:"column:address_id;type:uuid;not null"`
	VendorID  *uuid.UUID `gorm:"column:vendor_id;type:uuid"`
}

func (Warehouse) TableName() string {
	return "public.warehouses"
}
