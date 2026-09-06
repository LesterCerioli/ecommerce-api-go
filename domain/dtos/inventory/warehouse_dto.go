package inventory

import "github.com/google/uuid"

type WarehouseDTO struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	AddressID uuid.UUID  `json:"address_id"`
	VendorID  *uuid.UUID `json:"vendor_id"`
	VendorName string    `json:"vendor_name,omitempty"`
}
