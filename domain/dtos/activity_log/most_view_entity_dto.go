package activity_log

type MostViewEntityDTO struct {
	EntityID    string `json:"entity_id"`
	EntityTypeID string `json:"entity_type_id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	ViewedCount int    `json:"viewed_count"`
}
