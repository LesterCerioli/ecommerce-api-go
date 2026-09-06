package core

import "github.com/google/uuid"

type Role struct {
	ID               uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	Name             string     `gorm:"column:name"`
	NormalizedName   string     `gorm:"column:normalized_name"`
	ConcurrencyStamp string     `gorm:"column:concurrency_stamp"`
	Users            []UserRole `gorm:"foreignKey:RoleID;references:ID"`
}

func (Role) TableName() string {
	return "public.roles"
}
