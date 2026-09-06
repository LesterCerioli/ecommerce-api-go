package news

import "github.com/google/uuid"

type NewsItemCategory struct {
	CategoryID uuid.UUID     `gorm:"primaryKey;column:category_id;type:uuid"`
	Category   *NewsCategory `gorm:"foreignKey:CategoryID;references:ID"`
	NewsItemID uuid.UUID     `gorm:"primaryKey;column:news_item_id;type:uuid"`
	NewsItem   *NewsItem     `gorm:"foreignKey:NewsItemID;references:ID"`
}

func (NewsItemCategory) TableName() string {
	return "public.news_item_categories"
}
