package create

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderItem struct {
	ProductId uuid.UUID
	Quantity  int
	Price     decimal.Decimal
	SubTotal  decimal.Decimal
}
