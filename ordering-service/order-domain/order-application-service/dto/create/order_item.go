package create

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderItem struct {
	ProductId uuid.UUID       `json:"product_id"`
	Quantity  int             `json:"quantity"`
	Price     decimal.Decimal `json:"price"`
	SubTotal  decimal.Decimal `json:"sub_total"`
}
