package core

import "time"

type WidgetDTO struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	ViewComponentName string   `json:"view_component_name"`
	CreateURL        string    `json:"create_url"`
	EditURL          string    `json:"edit_url"`
	CreatedOn        time.Time `json:"created_on"`
	IsPublished      bool      `json:"is_published"`
}
