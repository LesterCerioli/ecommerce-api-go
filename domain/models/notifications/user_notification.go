package notifications

import (
	"time"

	"github.com/google/uuid"

	sharednotifications "ecommerce-api-go/domain/shared/domain/notifications"
)

type UserNotificationState = sharednotifications.UserNotificationState

type UserNotification struct {
	ID                   uuid.UUID             `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	UserID               uuid.UUID             `gorm:"column:user_id;type:uuid;not null"`
	NotificationDetailID uuid.UUID             `gorm:"column:notification_detail_id;type:uuid;not null"`
	NotificationDetail   *NotificationDetail   `gorm:"foreignKey:NotificationDetailID;references:ID"`
	State                UserNotificationState `gorm:"column:state;not null"`
	CreatedOn            time.Time             `gorm:"column:created_on;not null"`
	IsDeleted            bool                  `gorm:"column:is_deleted;not null"`
}

func (UserNotification) TableName() string {
	return "public.user_notifications"
}
