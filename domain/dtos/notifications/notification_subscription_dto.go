package notifications

import (
	"time"

	"github.com/google/uuid"
)

type NotificationSubscriptionDTO struct {
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"user_id"`
	NotificationName string  `json:"notification_name"`
	EntityTypeName string    `json:"entity_type_name"`
	EntityTypeAssemblyQualifiedName string `json:"entity_type_assembly_qualified_name"`
	EntityID       string    `json:"entity_id"`
	CreatedOn      time.Time `json:"created_on"`
	IsDeleted      bool      `json:"is_deleted"`
}
