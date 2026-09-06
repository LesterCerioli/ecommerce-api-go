package activity_log

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ID             uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	ActivityTypeID uuid.UUID  `gorm:"column:activity_type_id;type:uuid;not null"`
	ActivityType   *ActivityType `gorm:"foreignKey:ActivityTypeID;references:ID"`
	UserID         uuid.UUID  `gorm:"column:user_id;type:uuid;not null"`
	CreatedOn      time.Time  `gorm:"column:created_on;not null"`
	EntityID       uuid.UUID  `gorm:"column:entity_id;type:uuid;not null"`
	EntityTypeID   string     `gorm:"column:entity_type_id;size:450;not null"`
}

func (Activity) TableName() string {
	return "public.activities"
}
