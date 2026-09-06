package activity_log

import "github.com/google/uuid"

type ActivityType struct {
	ID   uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name string    `gorm:"column:name;size:450;not null"`
}

func (ActivityType) TableName() string {
	return "public.activity_types"
}
