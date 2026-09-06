package inventory

import (
	"time"

	"github.com/google/uuid"
)

type StockHistoryDTO struct {
	ID               uuid.UUID `json:"id"`
	ProductID        uuid.UUID `json:"product_id"`
	ProductName      string    `json:"product_name,omitempty"`
	WarehouseID      uuid.UUID `json:"warehouse_id"`
	WarehouseName    string    `json:"warehouse_name,omitempty"`
	CreatedOn        time.Time `json:"created_on"`
	CreatedByID      uuid.UUID `json:"created_by_id"`
	CreatedByName    string    `json:"created_by_name,omitempty"`
	AdjustedQuantity int64     `json:"adjusted_quantity"`
	Note             string    `json:"note"`
}
