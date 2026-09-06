package inventory

import "github.com/google/uuid"

type StockDTO struct {
	ID              uuid.UUID `json:"id"`
	ProductID       uuid.UUID `json:"product_id"`
	ProductName     string    `json:"product_name,omitempty"`
	WarehouseID     uuid.UUID `json:"warehouse_id"`
	WarehouseName   string    `json:"warehouse_name,omitempty"`
	Quantity        int       `json:"quantity"`
	ReservedQuantity int      `json:"reserved_quantity"`
}
