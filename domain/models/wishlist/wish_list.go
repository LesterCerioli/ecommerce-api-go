package wishlist

import (
	"time"

	"github.com/google/uuid"

	"ecommerce-api-go/domain/models/core"
)

type WishList struct {
	ID              uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	UserID          uuid.UUID      `gorm:"column:user_id;type:uuid;not null"`
	User            *core.User     `gorm:"foreignKey:UserID;references:ID"`
	SharingCode     string         `gorm:"column:sharing_code;size:450"`
	Items           []WishListItem `gorm:"foreignKey:WishListID;references:ID"`
	CreatedOn       time.Time      `gorm:"column:created_on;not null"`
	LatestUpdatedOn time.Time      `gorm:"column:latest_updated_on;not null"`
}

func (WishList) TableName() string {
	return "public.wish_lists"
}
