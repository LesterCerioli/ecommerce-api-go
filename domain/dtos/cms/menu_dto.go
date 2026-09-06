package cms

import "github.com/google/uuid"

type MenuDTO struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	IsPublished bool         `json:"is_published"`
	IsSystem    bool         `json:"is_system"`
	MenuItems   []MenuItemDTO `json:"menu_items,omitempty"`
}
