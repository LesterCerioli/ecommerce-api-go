package reviews

import (
	"time"

	"github.com/google/uuid"

	sharedreviews "ecommerce-api-go/domain/shared/domain/reviews"
	"ecommerce-api-go/domain/models/core"
)

type ReviewStatus = sharedreviews.ReviewStatus

type Review struct {
	ID           uuid.UUID     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	UserID       uuid.UUID     `gorm:"column:user_id;type:uuid;not null"`
	User         *core.User    `gorm:"foreignKey:UserID;references:ID"`
	Title        string        `gorm:"column:title;size:450"`
	Comment      string        `gorm:"column:comment"`
	Rating       int           `gorm:"column:rating;not null"`
	ReviewerName string        `gorm:"column:reviewer_name;size:450"`
	Status       ReviewStatus  `gorm:"column:status;not null"`
	CreatedOn    time.Time     `gorm:"column:created_on;not null"`
	EntityTypeID string        `gorm:"column:entity_type_id;size:450"`
	EntityID     uuid.UUID     `gorm:"column:entity_id;type:uuid;not null"`
	Replies      []Reply       `gorm:"foreignKey:ReviewID;references:ID"`
}

func (Review) TableName() string {
	return "public.reviews"
}
