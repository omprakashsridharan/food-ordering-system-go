package message

import (
	"common-domain/valueobject"
	"github.com/shopspring/decimal"
	"time"
)

type PaymentResponse struct {
	Id              string
	SagaId          string
	OrderId         string
	PaymentId       string
	CustomerId      string
	Price           decimal.Decimal
	CreatedAt       time.Time
	PaymentStatus   valueobject.PaymentStatus
	FailureMessages []string
}
