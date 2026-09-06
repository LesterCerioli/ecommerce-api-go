package core

import (
	"time"

	"github.com/google/uuid"
)

type WidgetInstance struct {
	ID              uuid.UUID   `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name            string      `gorm:"column:name"`
	CreatedOn       time.Time   `gorm:"column:created_on;not null"`
	LatestUpdatedOn time.Time   `gorm:"column:latest_updated_on;not null"`
	PublishStart    *time.Time  `gorm:"column:publish_start"`
	PublishEnd      *time.Time  `gorm:"column:publish_end"`
	WidgetID        string      `gorm:"column:widget_id"`
	Widget          *Widget     `gorm:"foreignKey:WidgetID;references:ID"`
	WidgetZoneID    uuid.UUID   `gorm:"column:widget_zone_id;type:uuid;not null"`
	WidgetZone      *WidgetZone `gorm:"foreignKey:WidgetZoneID;references:ID"`
	DisplayOrder    int         `gorm:"column:display_order;not null"`
	Data            string      `gorm:"column:data"`
	HtmlData        string      `gorm:"column:html_data"`
}

func (WidgetInstance) TableName() string {
	return "public.widget_instances"
}

func (w WidgetInstance) IsPublished() bool {
	now := time.Now()
	return w.PublishStart != nil && w.PublishStart.Before(now) && (w.PublishEnd == nil || w.PublishEnd.After(now))
}
