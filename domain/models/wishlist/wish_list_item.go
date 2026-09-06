package wishlist

import (
	"time"

	"github.com/google/uuid"
)

type WishListItem struct {
	ID              uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	WishListID      uuid.UUID `gorm:"column:wish_list_id;type:uuid;not null"`
	WishList        *WishList `gorm:"foreignKey:WishListID;references:ID"`
	ProductID       uuid.UUID `gorm:"column:product_id;type:uuid;not null"`
	Description     string    `gorm:"column:description"`
	Quantity        int       `gorm:"column:quantity;not null"`
	CreatedOn       time.Time `gorm:"column:created_on;not null"`
	LatestUpdatedOn time.Time `gorm:"column:latest_updated_on;not null"`
}

func (WishListItem) TableName() string {
	return "public.wish_list_items"
}
