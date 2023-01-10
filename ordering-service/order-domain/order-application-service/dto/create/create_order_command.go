package create

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateOrderCommand struct {
	CustomerId   uuid.UUID       `json:"customer_id"`
	RestaurantId uuid.UUID       `json:"restaurant_id"`
	Price        decimal.Decimal `json:"price"`
	Address      OrderAddress    `json:"address"`
	Items        []OrderItem     `json:"items"`
}
