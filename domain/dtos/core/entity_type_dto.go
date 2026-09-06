package core

type EntityTypeDTO struct {
	ID                string `json:"id"`
	IsMenuable        bool   `json:"is_menuable"`
	AreaName          string `json:"area_name"`
	RoutingController string `json:"routing_controller"`
	RoutingAction     string `json:"routing_action"`
}
