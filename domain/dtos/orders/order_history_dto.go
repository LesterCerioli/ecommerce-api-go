package orders

import (
	"time"

	"github.com/google/uuid"
)

type OrderHistoryDTO struct {
	ID            uuid.UUID  `json:"id"`
	OrderID       uuid.UUID  `json:"order_id"`
	OldStatus     *int       `json:"old_status"`
	NewStatus     int        `json:"new_status"`
	OrderSnapshot string     `json:"order_snapshot"`
	Note          string     `json:"note"`
	CreatedOn     time.Time  `json:"created_on"`
	CreatedByID   uuid.UUID  `json:"created_by_id"`
	CreatedByName string     `json:"created_by_name,omitempty"`
}
