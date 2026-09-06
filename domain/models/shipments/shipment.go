package shipments

import (
	"time"

	"github.com/google/uuid"

	"ecommerce-api-go/domain/models/core"
	"ecommerce-api-go/domain/models/inventory"
	"ecommerce-api-go/domain/models/orders"
)

type Shipment struct {
	ID              uuid.UUID       `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	OrderID         uuid.UUID       `gorm:"column:order_id;type:uuid;not null"`
	Order           *orders.Order   `gorm:"foreignKey:OrderID;references:ID"`
	TrackingNumber  string          `gorm:"column:tracking_number;size:450"`
	WarehouseID     uuid.UUID       `gorm:"column:warehouse_id;type:uuid;not null"`
	Warehouse       *inventory.Warehouse `gorm:"foreignKey:WarehouseID;references:ID"`
	VendorID        *uuid.UUID      `gorm:"column:vendor_id;type:uuid"`
	CreatedByID     uuid.UUID       `gorm:"column:created_by_id;type:uuid;not null"`
	CreatedBy       *core.User      `gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedOn       time.Time       `gorm:"column:created_on;not null"`
	LatestUpdatedOn time.Time       `gorm:"column:latest_updated_on;not null"`
	Items           []ShipmentItem  `gorm:"foreignKey:ShipmentID;references:ID"`
}

func (Shipment) TableName() string {
	return "public.shipments"
}
