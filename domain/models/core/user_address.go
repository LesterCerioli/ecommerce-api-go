package core

import (
	"time"

	"github.com/google/uuid"
)

type UserAddress struct {
	ID         uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	UserID     uuid.UUID  `gorm:"column:user_id;type:uuid;not null"`
	User       *User      `gorm:"foreignKey:UserID;references:ID"`
	AddressID  uuid.UUID  `gorm:"column:address_id;type:uuid;not null"`
	Address    *Address   `gorm:"foreignKey:AddressID;references:ID"`
	AddressType AddressType `gorm:"column:address_type;not null"`
	LastUsedOn *time.Time `gorm:"column:last_used_on"`
}

func (UserAddress) TableName() string {
	return "public.user_addresses"
}
