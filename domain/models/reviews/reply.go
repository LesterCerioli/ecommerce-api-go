package reviews

import (
	"time"

	"github.com/google/uuid"

	sharedreviews "ecommerce-api-go/domain/shared/domain/reviews"
	"ecommerce-api-go/domain/models/core"
)

type ReplyStatus = sharedreviews.ReplyStatus

type Reply struct {
	ID           uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	ReviewID     uuid.UUID  `gorm:"column:review_id;type:uuid;not null"`
	Review       *Review    `gorm:"foreignKey:ReviewID;references:ID"`
	UserID       uuid.UUID  `gorm:"column:user_id;type:uuid;not null"`
	User         *core.User `gorm:"foreignKey:UserID;references:ID"`
	Comment      string     `gorm:"column:comment"`
	ReplierName  string     `gorm:"column:replier_name;size:450"`
	Status       ReplyStatus `gorm:"column:status;not null"`
	CreatedOn    time.Time  `gorm:"column:created_on;not null"`
}

func (Reply) TableName() string {
	return "public.replies"
}
