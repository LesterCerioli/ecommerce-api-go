package product_recently_viewed

import (
	"time"

	"github.com/google/uuid"
)

type RecentlyViewedProductDTO struct {
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"user_id"`
	ProductID      uuid.UUID `json:"product_id"`
	LatestViewedOn time.Time `json:"latest_viewed_on"`
}
