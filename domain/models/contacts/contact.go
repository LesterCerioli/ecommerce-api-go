package contacts

import (
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	ID            uuid.UUID     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	FullName      string        `gorm:"column:full_name;size:450"`
	PhoneNumber   string        `gorm:"column:phone_number;size:450"`
	EmailAddress  string        `gorm:"column:email_address;size:450"`
	Address       string        `gorm:"column:address;size:450"`
	Content       string        `gorm:"column:content"`
	ContactAreaID uuid.UUID     `gorm:"column:contact_area_id;type:uuid;not null"`
	ContactArea   *ContactArea  `gorm:"foreignKey:ContactAreaID;references:ID"`
	IsDeleted     bool          `gorm:"column:is_deleted;not null"`
	CreatedOn     time.Time     `gorm:"column:created_on;not null"`
}

func (Contact) TableName() string {
	return "public.contacts"
}
