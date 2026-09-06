package shipments

type ShipmentStatus int

const (
)

func (s ShipmentStatus) String() string {
	return "Unknown"
}

func (s ShipmentStatus) IsValid() bool {
	return false
}

func ParseShipmentStatus(v int) (ShipmentStatus, bool) {
	s := ShipmentStatus(v)
	return s, s.IsValid()
}
