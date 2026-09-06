package notifications

import "time"

type NotificationDetailDTO struct {
	ID             string `json:"id"`
	NotificationName string `json:"notification_name"`
	Data           string `json:"data"`
	DataTypeName   string `json:"data_type_name"`
	EntityTypeName string `json:"entity_type_name"`
	EntityTypeAssemblyQualifiedName string `json:"entity_type_assembly_qualified_name"`
	EntityID       string `json:"entity_id"`
	Severity       int    `json:"severity"`
	CreatedOn      time.Time `json:"created_on"`
	IsDeleted      bool   `json:"is_deleted"`
}
