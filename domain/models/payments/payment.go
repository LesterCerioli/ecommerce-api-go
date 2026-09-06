package payments

import (
	"time"

	"github.com/google/uuid"

	sharedpayments "ecommerce-api-go/domain/shared/domain/payments"
	"ecommerce-api-go/domain/models/orders"
)

type PaymentStatus = sharedpayments.PaymentStatus

type Payment struct {
	ID                   uuid.UUID       `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();<-:false"`
	OrderID              uuid.UUID       `gorm:"column:order_id;type:uuid;not null"`
	Order                *orders.Order   `gorm:"foreignKey:OrderID;references:ID"`
	CreatedOn            time.Time       `gorm:"column:created_on;not null"`
	LatestUpdatedOn      time.Time       `gorm:"column:latest_updated_on;not null"`
	Amount               float64         `gorm:"column:amount;not null"`
	PaymentFee           float64         `gorm:"column:payment_fee;not null"`
	PaymentMethod        string          `gorm:"column:payment_method;size:450"`
	GatewayTransactionID string          `gorm:"column:gateway_transaction_id;size:450"`
	Status               PaymentStatus   `gorm:"column:status;not null"`
	FailureMessage       string          `gorm:"column:failure_message"`
}

func (Payment) TableName() string {
	return "public.payments"
}
