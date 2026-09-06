package core

import "github.com/google/uuid"

type UserRole struct {
	UserID uuid.UUID `gorm:"primaryKey;column:user_id;type:uuid"`
	User   *User     `gorm:"foreignKey:UserID;references:ID"`
	RoleID uuid.UUID `gorm:"primaryKey;column:role_id;type:uuid"`
	Role   *Role     `gorm:"foreignKey:RoleID;references:ID"`
}

func (UserRole) TableName() string {
	return "public.user_roles"
}
