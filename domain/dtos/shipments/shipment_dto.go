package shipments

import (
	"time"

	"github.com/google/uuid"
)

type ShipmentDTO struct {
	ID             uuid.UUID       `json:"id"`
	OrderID        uuid.UUID       `json:"order_id"`
	TrackingNumber string          `json:"tracking_number"`
	WarehouseID    uuid.UUID       `json:"warehouse_id"`
	WarehouseName  string          `json:"warehouse_name,omitempty"`
	VendorID       *uuid.UUID      `json:"vendor_id"`
	CreatedByID    uuid.UUID       `json:"created_by_id"`
	CreatedByName  string          `json:"created_by_name,omitempty"`
	CreatedOn      time.Time       `json:"created_on"`
	LatestUpdatedOn time.Time      `json:"latest_updated_on"`
	Items          []ShipmentItemDTO `json:"items,omitempty"`
}
