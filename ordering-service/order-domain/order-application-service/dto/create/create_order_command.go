package create

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateOrderCommand struct {
	CustomerId   uuid.UUID
	RestaurantId uuid.UUID
	Price        decimal.Decimal
	Address      OrderAddress
	Items        []OrderItem
}
