package shipments

import "github.com/google/uuid"

type ShipmentItem struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	ShipmentID    uuid.UUID `gorm:"column:shipment_id;type:uuid;not null"`
	Shipment      *Shipment `gorm:"foreignKey:ShipmentID;references:ID"`
	OrderItemID   uuid.UUID `gorm:"column:order_item_id;type:uuid;not null"`
	ProductID     uuid.UUID `gorm:"column:product_id;type:uuid;not null"`
	Quantity      int       `gorm:"column:quantity;not null"`
}

func (ShipmentItem) TableName() string {
	return "public.shipment_items"
}
