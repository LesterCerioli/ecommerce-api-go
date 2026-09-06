package core

import "github.com/google/uuid"

type StateOrProvince struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	CountryID string    `gorm:"column:country_id"`
	Country   *Country  `gorm:"foreignKey:CountryID;references:ID"`
	Code      string    `gorm:"column:code"`
	Name      string    `gorm:"column:name;not null"`
	Type      string    `gorm:"column:type"`
}

func (StateOrProvince) TableName() string {
	return "public.states_or_provinces"
}
