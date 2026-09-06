package payments

type PaymentStatus int

const (
	PaymentStatusSucceeded PaymentStatus = 1
	PaymentStatusFailed    PaymentStatus = 5
)

func (s PaymentStatus) String() string {
	switch s {
	case PaymentStatusSucceeded:
		return "Succeeded"
	case PaymentStatusFailed:
		return "Failed"
	default:
		return "Unknown"
	}
}

func (s PaymentStatus) IsValid() bool {
	switch s {
	case PaymentStatusSucceeded, PaymentStatusFailed:
		return true
	default:
		return false
	}
}

func ParsePaymentStatus(v int) (PaymentStatus, bool) {
	s := PaymentStatus(v)
	return s, s.IsValid()
}
