package catalog

import (
	"time"

	"github.com/google/uuid"
)

type ProductDTO struct {
	ID                    uuid.UUID   `json:"id"`
	Name                  string      `json:"name"`
	Slug                  string      `json:"slug"`
	MetaTitle             string      `json:"meta_title"`
	MetaKeywords          string      `json:"meta_keywords"`
	MetaDescription       string      `json:"meta_description"`
	IsPublished           bool        `json:"is_published"`
	PublishedOn           *time.Time  `json:"published_on"`
	IsDeleted             bool        `json:"is_deleted"`
	CreatedByID           uuid.UUID   `json:"created_by_id"`
	CreatedByName         string      `json:"created_by_name,omitempty"`
	CreatedOn             time.Time   `json:"created_on"`
	LatestUpdatedOn       time.Time   `json:"latest_updated_on"`
	LatestUpdatedByID     uuid.UUID   `json:"latest_updated_by_id"`
	LatestUpdatedByName   string      `json:"latest_updated_by_name,omitempty"`
	ShortDescription      string      `json:"short_description"`
	Description           string      `json:"description"`
	Specification         string      `json:"specification"`
	Price                 float64     `json:"price"`
	OldPrice              *float64    `json:"old_price"`
	SpecialPrice          *float64    `json:"special_price"`
	SpecialPriceStart     *time.Time  `json:"special_price_start"`
	SpecialPriceEnd       *time.Time  `json:"special_price_end"`
	HasOptions            bool        `json:"has_options"`
	IsVisibleIndividually bool        `json:"is_visible_individually"`
	IsFeatured            bool        `json:"is_featured"`
	IsCallForPricing      bool        `json:"is_call_for_pricing"`
	IsAllowToOrder        bool        `json:"is_allow_to_order"`
	StockTrackingIsEnabled bool       `json:"stock_tracking_is_enabled"`
	StockQuantity         int         `json:"stock_quantity"`
	Sku                   string      `json:"sku"`
	Gtin                  string      `json:"gtin"`
	NormalizedName        string      `json:"normalized_name"`
	DisplayOrder          int         `json:"display_order"`
	VendorID              *uuid.UUID  `json:"vendor_id"`
	VendorName            string      `json:"vendor_name,omitempty"`
	ThumbnailImageID      *uuid.UUID  `json:"thumbnail_image_id"`
	ReviewsCount          int         `json:"reviews_count"`
	RatingAverage         *float64    `json:"rating_average"`
	BrandID               *uuid.UUID  `json:"brand_id"`
	BrandName             string      `json:"brand_name,omitempty"`
	TaxClassID            *uuid.UUID  `json:"tax_class_id"`
	TaxClassName          string      `json:"tax_class_name,omitempty"`
}
