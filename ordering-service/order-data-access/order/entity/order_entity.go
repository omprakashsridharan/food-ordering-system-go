package entity

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OrderEntity struct {
	gorm.Model
	Id              uuid.UUID
	CustomerId      uuid.UUID
	RestaurantId    uuid.UUID
	TrackingId      uuid.UUID
	Price           decimal.Decimal
	OrderStatus     string
	FailureMessages string
	OrderAddress    OrderAddressEntity `gorm:"foreignKey:OrderId;references:id"`
	OrderItems      []OrderItemEntity  `gorm:"foreignKey:OrderId;references:id"`
}
