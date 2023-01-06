package valueobject

type OrderStatus string

const (
	ORDER_STATUS_PENDING    OrderStatus = "pending"
	ORDER_STATUS_PAID                   = "rejected"
	ORDER_STATUS_APPROVED               = "approved"
	ORDER_STATUS_CANCELLING             = "cancelling"
	ORDER_STATUS_CANCELLED              = "cancelled"
)
