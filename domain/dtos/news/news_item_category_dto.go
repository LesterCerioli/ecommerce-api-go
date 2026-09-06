package news

import "github.com/google/uuid"

type NewsItemCategoryDTO struct {
	CategoryID uuid.UUID `json:"category_id"`
	NewsItemID uuid.UUID `json:"news_item_id"`
}
