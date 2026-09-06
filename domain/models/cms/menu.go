package cms

import "github.com/google/uuid"

type Menu struct {
	ID         uuid.UUID    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name       string       `gorm:"column:name;size:450;not null"`
	IsPublished bool        `gorm:"column:is_published;not null"`
	IsSystem   bool         `gorm:"column:is_system;not null"`
	MenuItems  []MenuItem   `gorm:"foreignKey:MenuID;references:ID"`
}

func (Menu) TableName() string {
	return "public.menus"
}
