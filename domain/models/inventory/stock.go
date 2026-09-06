package inventory

import "github.com/google/uuid"

type Stock struct {
	ID              uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	ProductID       uuid.UUID `gorm:"column:product_id;type:uuid;not null"`
	WarehouseID     uuid.UUID `gorm:"column:warehouse_id;type:uuid;not null"`
	Warehouse       *Warehouse `gorm:"foreignKey:WarehouseID;references:ID"`
	Quantity        int       `gorm:"column:quantity;not null"`
	ReservedQuantity int      `gorm:"column:reserved_quantity;not null"`
}

func (Stock) TableName() string {
	return "public.stocks"
}
