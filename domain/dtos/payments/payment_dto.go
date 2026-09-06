package payments

import (
	"time"

	"github.com/google/uuid"
)

type PaymentDTO struct {
	ID                   uuid.UUID `json:"id"`
	OrderID              uuid.UUID `json:"order_id"`
	CreatedOn            time.Time `json:"created_on"`
	LatestUpdatedOn      time.Time `json:"latest_updated_on"`
	Amount               float64   `json:"amount"`
	PaymentFee           float64   `json:"payment_fee"`
	PaymentMethod        string    `json:"payment_method"`
	GatewayTransactionID string    `json:"gateway_transaction_id"`
	Status               int       `json:"status"`
	FailureMessage       string    `json:"failure_message"`
}
