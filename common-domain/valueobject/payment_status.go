package valueobject

type PaymentStatus string

const (
	PAYMENT_STATUS_COMPLETED PaymentStatus = "completed"
	PAYMENT_STATUS_FAILED                  = "failed"
	PAYMENT_STATUS_CANCELLED               = "cancelled"
)
