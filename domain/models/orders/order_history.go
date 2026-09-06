package orders

import (
	"time"

	"github.com/google/uuid"

	"ecommerce-api-go/domain/models/core"
)

type OrderHistory struct {
	ID            uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	OrderID       uuid.UUID  `gorm:"column:order_id;type:uuid;not null"`
	Order         *Order     `gorm:"foreignKey:OrderID;references:ID"`
	OldStatus     *OrderStatus `gorm:"column:old_status"`
	NewStatus     OrderStatus `gorm:"column:new_status;not null"`
	OrderSnapshot string     `gorm:"column:order_snapshot"`
	Note          string     `gorm:"column:note;size:1000"`
	CreatedOn     time.Time  `gorm:"column:created_on;not null"`
	CreatedByID   uuid.UUID  `gorm:"column:created_by_id;type:uuid;not null"`
	CreatedBy     *core.User `gorm:"foreignKey:CreatedByID;references:ID"`
}

func (OrderHistory) TableName() string {
	return "public.order_histories"
}
