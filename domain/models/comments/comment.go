package comments

import (
	"github.com/google/uuid"

	sharedcomments "ecommerce-api-go/domain/shared/domain/comments"
)

// Alias local: o tipo mora no pacote shared (que também se chama
// "comments"), por isso o uso sem qualificação quebra.
// Com o alias, `CommentStatus` aqui equivale a `sharedcomments.CommentStatus`.
type CommentStatus = sharedcomments.CommentStatus

type Comment struct {
	ID            uuid.UUID     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CommentText   string        `gorm:"column:comment_text;not null"`
	UserID        uuid.UUID     `gorm:"column:user_id;type:uuid;not null"`
	CommenterName string        `gorm:"column:commenter_name;not null"`
	Status        CommentStatus `gorm:"column:status;not null"`
	EntityID      uuid.UUID     `gorm:"column:entity_id;type:uuid;not null"`
}

func (Comment) TableName() string {
	return "public.comments"
}
