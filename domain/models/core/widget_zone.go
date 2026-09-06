package core

import "github.com/google/uuid"

type WidgetZone struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name        string    `gorm:"column:name;not null"`
	Description string    `gorm:"column:description"`
}

func (WidgetZone) TableName() string {
	return "public.widget_zones"
}
