package cms

import "github.com/google/uuid"

type MenuItem struct {
	ID           uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	ParentID     *uuid.UUID `gorm:"column:parent_id;type:uuid"`
	Parent       *MenuItem  `gorm:"foreignKey:ParentID;references:ID"`
	Children     []MenuItem `gorm:"foreignKey:ParentID;references:ID"`
	MenuID       uuid.UUID  `gorm:"column:menu_id;type:uuid;not null"`
	Menu         *Menu      `gorm:"foreignKey:MenuID;references:ID"`
	EntityID     *uuid.UUID `gorm:"column:entity_id;type:uuid"`
	CustomLink   string     `gorm:"column:custom_link;size:450"`
	Name         string     `gorm:"column:name;size:450"`
	DisplayOrder int        `gorm:"column:display_order;not null"`
}

func (MenuItem) TableName() string {
	return "public.menu_items"
}
