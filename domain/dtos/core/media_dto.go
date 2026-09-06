package core

import "github.com/google/uuid"

type MediaDTO struct {
	ID        uuid.UUID `json:"id"`
	Caption   string    `json:"caption"`
	FileSize  int       `json:"file_size"`
	FileName  string    `json:"file_name"`
	MediaType int       `json:"media_type"`
}
