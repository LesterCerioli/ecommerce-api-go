package shipments

import "github.com/google/uuid"

type ShipmentItemDTO struct {
	ID          uuid.UUID `json:"id"`
	ShipmentID  uuid.UUID `json:"shipment_id"`
	OrderItemID uuid.UUID `json:"order_item_id"`
	ProductID   uuid.UUID `json:"product_id"`
	ProductName string    `json:"product_name,omitempty"`
	Quantity    int       `json:"quantity"`
}
