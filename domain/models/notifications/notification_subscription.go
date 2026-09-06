package notifications

import (
	"time"

	"github.com/google/uuid"
)

type NotificationSubscription struct {
	ID             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	UserID         uuid.UUID `gorm:"column:user_id;type:uuid;not null"`
	NotificationName string  `gorm:"column:notification_name;size:96"`
	EntityTypeName string    `gorm:"column:entity_type_name;size:250"`
	EntityTypeAssemblyQualifiedName string `gorm:"column:entity_type_assembly_qualified_name;size:512"`
	EntityID       string    `gorm:"column:entity_id;size:96"`
	CreatedOn      time.Time `gorm:"column:created_on;not null"`
	IsDeleted      bool      `gorm:"column:is_deleted;not null"`
}

func (NotificationSubscription) TableName() string {
	return "public.notification_subscriptions"
}
