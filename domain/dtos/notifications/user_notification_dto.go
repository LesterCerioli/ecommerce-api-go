package notifications

import (
	"time"

	"github.com/google/uuid"
)

type UserNotificationDTO struct {
	ID                   uuid.UUID `json:"id"`
	UserID               uuid.UUID `json:"user_id"`
	NotificationDetailID uuid.UUID `json:"notification_detail_id"`
	State                int       `json:"state"`
	CreatedOn            time.Time `json:"created_on"`
	IsDeleted            bool      `json:"is_deleted"`
	Detail               *NotificationDetailDTO `json:"detail,omitempty"`
}
