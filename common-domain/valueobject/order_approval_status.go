package valueobject

type OrderApprovalStatus string

const (
	ORDER_APPROVAL_STATUS_APPROVED OrderApprovalStatus = "approved"
	ORDER_APPROVAL_STATUS_REJECTED                     = "rejected"
)
