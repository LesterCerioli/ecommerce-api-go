package catalog

import (
	"time"

	"github.com/google/uuid"

	"ecommerce-api-go/domain/models/core"
	"ecommerce-api-go/domain/models/tax"
)

type Product struct {
	ID                    uuid.UUID              `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name                  string                 `gorm:"column:name;not null"`
	Slug                  string                 `gorm:"column:slug;not null"`
	MetaTitle             string                 `gorm:"column:meta_title"`
	MetaKeywords          string                 `gorm:"column:meta_keywords"`
	MetaDescription       string                 `gorm:"column:meta_description"`
	IsPublished           bool                   `gorm:"column:is_published;not null"`
	PublishedOn           *time.Time             `gorm:"column:published_on"`
	IsDeleted             bool                   `gorm:"column:is_deleted;not null"`
	CreatedByID           uuid.UUID              `gorm:"column:created_by_id;type:uuid;not null"`
	CreatedBy             *core.User             `gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedOn             time.Time              `gorm:"column:created_on;not null"`
	LatestUpdatedOn       time.Time              `gorm:"column:latest_updated_on;not null"`
	LatestUpdatedByID     uuid.UUID              `gorm:"column:latest_updated_by_id;type:uuid;not null"`
	LatestUpdatedBy       *core.User             `gorm:"foreignKey:LatestUpdatedByID;references:ID"`
	ShortDescription      string                 `gorm:"column:short_description;size:450"`
	Description           string                 `gorm:"column:description"`
	Specification         string                 `gorm:"column:specification"`
	Price                 float64                `gorm:"column:price;not null"`
	OldPrice              *float64               `gorm:"column:old_price"`
	SpecialPrice          *float64               `gorm:"column:special_price"`
	SpecialPriceStart     *time.Time             `gorm:"column:special_price_start"`
	SpecialPriceEnd       *time.Time             `gorm:"column:special_price_end"`
	HasOptions            bool                   `gorm:"column:has_options;not null"`
	IsVisibleIndividually bool                   `gorm:"column:is_visible_individually;not null"`
	IsFeatured            bool                   `gorm:"column:is_featured;not null"`
	IsCallForPricing      bool                   `gorm:"column:is_call_for_pricing;not null"`
	IsAllowToOrder        bool                   `gorm:"column:is_allow_to_order;not null"`
	StockTrackingIsEnabled bool                  `gorm:"column:stock_tracking_is_enabled;not null"`
	StockQuantity         int                    `gorm:"column:stock_quantity;not null"`
	Sku                   string                 `gorm:"column:sku;size:450"`
	Gtin                  string                 `gorm:"column:gtin;size:450"`
	NormalizedName        string                 `gorm:"column:normalized_name;size:450"`
	DisplayOrder          int                    `gorm:"column:display_order;not null"`
	VendorID              *uuid.UUID             `gorm:"column:vendor_id;type:uuid"`
	Vendor                *core.Vendor           `gorm:"foreignKey:VendorID;references:ID"`
	ThumbnailImageID      *uuid.UUID             `gorm:"column:thumbnail_image_id;type:uuid"`
	ThumbnailImage        *core.Media            `gorm:"foreignKey:ThumbnailImageID;references:ID"`
	ReviewsCount          int                    `gorm:"column:reviews_count;not null"`
	RatingAverage         *float64               `gorm:"column:rating_average"`
	BrandID               *uuid.UUID             `gorm:"column:brand_id;type:uuid"`
	Brand                 *Brand                 `gorm:"foreignKey:BrandID;references:ID"`
	TaxClassID            *uuid.UUID             `gorm:"column:tax_class_id;type:uuid"`
	TaxClass              *tax.TaxClass          `gorm:"foreignKey:TaxClassID;references:ID"`
	Medias                []ProductMedia         `gorm:"foreignKey:ProductID;references:ID"`
	ProductLinks          []ProductLink          `gorm:"foreignKey:ProductID;references:ID"`
	LinkedProductLinks    []ProductLink          `gorm:"foreignKey:LinkedProductID;references:ID"`
	AttributeValues       []ProductAttributeValue `gorm:"foreignKey:ProductID;references:ID"`
	OptionValues          []ProductOptionValue   `gorm:"foreignKey:ProductID;references:ID"`
	Categories            []ProductCategory      `gorm:"foreignKey:ProductID;references:ID"`
	PriceHistories        []ProductPriceHistory  `gorm:"foreignKey:ProductID;references:ID"`
	OptionCombinations    []ProductOptionCombination `gorm:"foreignKey:ProductID;references:ID"`
}

func (Product) TableName() string {
	return "public.products"
}
