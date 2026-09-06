package core

import (
	"time"

	"github.com/google/uuid"
)

type WidgetInstanceDTO struct {
	ID              uuid.UUID  `json:"id"`
	Name            string     `json:"name"`
	CreatedOn       time.Time  `json:"created_on"`
	LatestUpdatedOn time.Time  `json:"latest_updated_on"`
	PublishStart    *time.Time `json:"publish_start"`
	PublishEnd      *time.Time `json:"publish_end"`
	WidgetID        string     `json:"widget_id"`
	WidgetName      string     `json:"widget_name,omitempty"`
	WidgetZoneID    uuid.UUID  `json:"widget_zone_id"`
	WidgetZoneName  string     `json:"widget_zone_name,omitempty"`
	DisplayOrder    int        `json:"display_order"`
	Data            string     `json:"data"`
	HtmlData        string     `json:"html_data"`
}
