package cms

import "github.com/google/uuid"

type MenuItemDTO struct {
	ID           uuid.UUID  `json:"id"`
	ParentID     *uuid.UUID `json:"parent_id"`
	MenuID       uuid.UUID  `json:"menu_id"`
	EntityID     *uuid.UUID `json:"entity_id"`
	CustomLink   string     `json:"custom_link"`
	Name         string     `json:"name"`
	DisplayOrder int        `json:"display_order"`
}
