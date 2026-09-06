package core

import "github.com/google/uuid"

type Entity struct {
	ID           uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Slug         string     `gorm:"column:slug;not null"`
	Name         string     `gorm:"column:name;not null"`
	EntityID     uuid.UUID  `gorm:"column:entity_id;type:uuid;not null"`
	EntityTypeID string     `gorm:"column:entity_type_id;not null"`
	EntityType   *EntityType `gorm:"foreignKey:EntityTypeID;references:ID"`
}

func (Entity) TableName() string {
	return "public.entities"
}
