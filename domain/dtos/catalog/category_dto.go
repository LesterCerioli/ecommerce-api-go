package catalog

import (
	"time"

	"github.com/google/uuid"
)

type CategoryDTO struct {
	ID               uuid.UUID  `json:"id"`
	Name             string     `json:"name"`
	Slug             string     `json:"slug"`
	MetaTitle        string     `json:"meta_title"`
	MetaKeywords     string     `json:"meta_keywords"`
	MetaDescription  string     `json:"meta_description"`
	Description      string     `json:"description"`
	DisplayOrder     int        `json:"display_order"`
	IsPublished      bool       `json:"is_published"`
	IncludeInMenu    bool       `json:"include_in_menu"`
	IsDeleted        bool       `json:"is_deleted"`
	ParentID         *uuid.UUID `json:"parent_id"`
	ParentName       string     `json:"parent_name,omitempty"`
	ThumbnailImageID *uuid.UUID `json:"thumbnail_image_id"`
	CreatedOn        time.Time  `json:"created_on"`
	LatestUpdatedOn  time.Time  `json:"latest_updated_on"`
}
