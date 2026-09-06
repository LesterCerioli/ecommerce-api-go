package notifications

import "time"

type NotificationSchemeDTO struct {
	ID             string `json:"id"`
	NotificationName string `json:"notification_name"`
	Data           string `json:"data"`
	DataTypeName   string `json:"data_type_name"`
	EntityTypeName string `json:"entity_type_name"`
	EntityTypeAssemblyQualifiedName string `json:"entity_type_assembly_qualified_name"`
	EntityID       string `json:"entity_id"`
	Severity       int    `json:"severity"`
	UserIds        string `json:"user_ids"`
	ExcludedUserIds string `json:"excluded_user_ids"`
	CreatedOn      time.Time `json:"created_on"`
	IsDeleted      bool   `json:"is_deleted"`
}
