package core

type EntityType struct {
	ID                string `gorm:"primaryKey;column:id;size:450"`
	IsMenuable        bool   `gorm:"column:is_menuable;not null"`
	AreaName          string `gorm:"column:area_name"`
	RoutingController string `gorm:"column:routing_controller"`
	RoutingAction     string `gorm:"column:routing_action"`
}

func (EntityType) TableName() string {
	return "public.entity_types"
}

func (e EntityType) Name() string {
	return e.ID
}
