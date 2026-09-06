package product_recently_viewed

import (
	"time"

	"github.com/google/uuid"
)

type RecentlyViewedProduct struct {
	ID             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	UserID         uuid.UUID `gorm:"column:user_id;type:uuid;not null"`
	ProductID      uuid.UUID `gorm:"column:product_id;type:uuid;not null"`
	LatestViewedOn time.Time `gorm:"column:latest_viewed_on;not null"`
}

func (RecentlyViewedProduct) TableName() string {
	return "public.recently_viewed_products"
}
