package message

import (
	"common-domain/valueobject"
	"time"
)

type RestaurantApprovalResponse struct {
	Id                  string
	SagaId              string
	OrderId             string
	RestaurantId        string
	CreatedAt           time.Time
	OrderApprovalStatus valueobject.OrderApprovalStatus
	FailureMessages     []string
}
