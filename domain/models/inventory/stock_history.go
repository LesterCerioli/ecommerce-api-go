package inventory

import (
	"time"

	"github.com/google/uuid"

	"ecommerce-api-go/domain/models/core"
)

type StockHistory struct {
	ID                uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	ProductID         uuid.UUID  `gorm:"column:product_id;type:uuid;not null"`
	WarehouseID       uuid.UUID  `gorm:"column:warehouse_id;type:uuid;not null"`
	Warehouse         *Warehouse `gorm:"foreignKey:WarehouseID;references:ID"`
	CreatedOn         time.Time  `gorm:"column:created_on;not null"`
	CreatedByID       uuid.UUID  `gorm:"column:created_by_id;type:uuid;not null"`
	CreatedBy         *core.User `gorm:"foreignKey:CreatedByID;references:ID"`
	AdjustedQuantity  int64      `gorm:"column:adjusted_quantity;not null"`
	Note              string     `gorm:"column:note;size:1000"`
}

func (StockHistory) TableName() string {
	return "public.stock_histories"
}
