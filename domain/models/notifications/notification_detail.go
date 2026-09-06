package notifications

import (
	"time"

	"github.com/google/uuid"
)

type NotificationDetail struct {
	ID             uuid.UUID            `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	NotificationName string             `gorm:"column:notification_name;size:96;not null"`
	Data           string               `gorm:"column:data"`
	DataTypeName   string               `gorm:"column:data_type_name;size:512"`
	EntityTypeName string               `gorm:"column:entity_type_name;size:250"`
	EntityTypeAssemblyQualifiedName string `gorm:"column:entity_type_assembly_qualified_name;size:512"`
	EntityID       string               `gorm:"column:entity_id;size:96"`
	Severity       NotificationSeverity `gorm:"column:severity;not null"`
	CreatedOn      time.Time            `gorm:"column:created_on;not null"`
	IsDeleted      bool                 `gorm:"column:is_deleted;not null"`
}

func (NotificationDetail) TableName() string {
	return "public.notification_details"
}
