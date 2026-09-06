package core

import "github.com/google/uuid"

type District struct {
	ID                uuid.UUID       `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	StateOrProvinceID uuid.UUID       `gorm:"column:state_or_province_id;type:uuid;not null"`
	StateOrProvince   *StateOrProvince `gorm:"foreignKey:StateOrProvinceID;references:ID"`
	Name              string          `gorm:"column:name;not null"`
	Type              string          `gorm:"column:type"`
	Location          string          `gorm:"column:location"`
}

func (District) TableName() string {
	return "public.districts"
}
