package core

import "github.com/google/uuid"

type Media struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Caption   string    `gorm:"column:caption"`
	FileSize  int       `gorm:"column:file_size;not null"`
	FileName  string    `gorm:"column:file_name"`
	MediaType MediaType `gorm:"column:media_type;not null"`
}

func (Media) TableName() string {
	return "public.media"
}
