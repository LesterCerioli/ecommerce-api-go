package contacts

import "github.com/google/uuid"

type ContactArea struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name      string    `gorm:"column:name;size:450;not null"`
	IsDeleted bool      `gorm:"column:is_deleted;not null"`
}

func (ContactArea) TableName() string {
	return "public.contact_areas"
}
