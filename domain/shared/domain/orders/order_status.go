package orders

type OrderStatus int

const (
	OrderStatusNew             OrderStatus = 1
	OrderStatusOnHold          OrderStatus = 10
	OrderStatusPendingPayment  OrderStatus = 20
	OrderStatusPaymentReceived OrderStatus = 30
	OrderStatusPaymentFailed   OrderStatus = 35
	OrderStatusInvoiced        OrderStatus = 40
	OrderStatusShipping        OrderStatus = 50
	OrderStatusShipped         OrderStatus = 60
	OrderStatusComplete        OrderStatus = 70
	OrderStatusCanceled        OrderStatus = 80
	OrderStatusRefunded        OrderStatus = 90
	OrderStatusClosed          OrderStatus = 100
)

func (s OrderStatus) String() string {
	switch s {
	case OrderStatusNew:
		return "New"
	case OrderStatusOnHold:
		return "OnHold"
	case OrderStatusPendingPayment:
		return "PendingPayment"
	case OrderStatusPaymentReceived:
		return "PaymentReceived"
	case OrderStatusPaymentFailed:
		return "PaymentFailed"
	case OrderStatusInvoiced:
		return "Invoiced"
	case OrderStatusShipping:
		return "Shipping"
	case OrderStatusShipped:
		return "Shipped"
	case OrderStatusComplete:
		return "Complete"
	case OrderStatusCanceled:
		return "Canceled"
	case OrderStatusRefunded:
		return "Refunded"
	case OrderStatusClosed:
		return "Closed"
	default:
		return "Unknown"
	}
}

func (s OrderStatus) IsValid() bool {
	switch s {
	case OrderStatusNew, OrderStatusOnHold, OrderStatusPendingPayment,
		OrderStatusPaymentReceived, OrderStatusPaymentFailed, OrderStatusInvoiced,
		OrderStatusShipping, OrderStatusShipped, OrderStatusComplete,
		OrderStatusCanceled, OrderStatusRefunded, OrderStatusClosed:
		return true
	default:
		return false
	}
}

func ParseOrderStatus(v int) (OrderStatus, bool) {
	s := OrderStatus(v)
	return s, s.IsValid()
}
