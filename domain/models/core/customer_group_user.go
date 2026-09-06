package core

import "github.com/google/uuid"

type CustomerGroupUser struct {
	UserID          uuid.UUID      `gorm:"primaryKey;column:user_id;type:uuid"`
	User            *User          `gorm:"foreignKey:UserID;references:ID"`
	CustomerGroupID uuid.UUID      `gorm:"primaryKey;column:customer_group_id;type:uuid"`
	CustomerGroup   *CustomerGroup `gorm:"foreignKey:CustomerGroupID;references:ID"`
}

func (CustomerGroupUser) TableName() string {
	return "public.customer_group_users"
}
