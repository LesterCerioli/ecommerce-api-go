package core

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                      uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	UserGUID                uuid.UUID  `gorm:"column:user_guid;type:uuid;not null"`
	UserName                string     `gorm:"column:user_name"`
	NormalizedUserName      string     `gorm:"column:normalized_user_name"`
	Email                   string     `gorm:"column:email"`
	NormalizedEmail         string     `gorm:"column:normalized_email"`
	EmailConfirmed          bool       `gorm:"column:email_confirmed;not null"`
	PasswordHash            string     `gorm:"column:password_hash"`
	SecurityStamp           string     `gorm:"column:security_stamp"`
	ConcurrencyStamp        string     `gorm:"column:concurrency_stamp"`
	PhoneNumber             string     `gorm:"column:phone_number"`
	PhoneNumberConfirmed    bool       `gorm:"column:phone_number_confirmed;not null"`
	TwoFactorEnabled        bool       `gorm:"column:two_factor_enabled;not null"`
	LockoutEnd              *time.Time `gorm:"column:lockout_end"`
	LockoutEnabled          bool       `gorm:"column:lockout_enabled;not null"`
	AccessFailedCount       int        `gorm:"column:access_failed_count;not null"`
	FullName                string     `gorm:"column:full_name;not null"`
	VendorID                *uuid.UUID `gorm:"column:vendor_id;type:uuid"`
	Vendor                  *Vendor    `gorm:"foreignKey:VendorID;references:ID"`
	IsDeleted               bool       `gorm:"column:is_deleted;not null"`
	CreatedOn               time.Time  `gorm:"column:created_on;not null"`
	LatestUpdatedOn         time.Time  `gorm:"column:latest_updated_on;not null"`
	UserAddresses           []UserAddress `gorm:"foreignKey:UserID;references:ID"`
	DefaultShippingAddressID *uuid.UUID   `gorm:"column:default_shipping_address_id;type:uuid"`
	DefaultShippingAddress   *UserAddress `gorm:"foreignKey:DefaultShippingAddressID;references:ID"`
	DefaultBillingAddressID  *uuid.UUID   `gorm:"column:default_billing_address_id;type:uuid"`
	DefaultBillingAddress    *UserAddress `gorm:"foreignKey:DefaultBillingAddressID;references:ID"`
	RefreshTokenHash        string     `gorm:"column:refresh_token_hash"`
	Roles                   []UserRole `gorm:"foreignKey:UserID;references:ID"`
	CustomerGroups          []CustomerGroupUser `gorm:"foreignKey:UserID;references:ID"`
	Culture                 string     `gorm:"column:culture"`
	ExtensionData           string     `gorm:"column:extension_data"`
}

func (User) TableName() string {
	return "public.users"
}
