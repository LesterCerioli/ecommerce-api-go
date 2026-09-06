package core

import "time"

type Widget struct {
	ID               string    `gorm:"primaryKey;column:id;size:450"`
	Name             string    `gorm:"column:name;not null"`
	ViewComponentName string   `gorm:"column:view_component_name"`
	CreateURL        string    `gorm:"column:create_url"`
	EditURL          string    `gorm:"column:edit_url"`
	CreatedOn        time.Time `gorm:"column:created_on;not null"`
	IsPublished      bool      `gorm:"column:is_published;not null"`
	Instances        []WidgetInstance `gorm:"foreignKey:WidgetID;references:ID"`
}

func (Widget) TableName() string {
	return "public.widgets"
}

func (w Widget) Code() string {
	return w.ID
}
