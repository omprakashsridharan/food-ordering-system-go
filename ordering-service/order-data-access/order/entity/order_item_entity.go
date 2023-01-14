package entity

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OrderItemEntity struct {
	gorm.Model
	Id        int64     `gorm:"primaryKey"`
	OrderId   uuid.UUID `gorm:"primaryKey"`
	ProductId uuid.UUID
	Price     decimal.Decimal
	Quantity  int64
	SubTotal  decimal.Decimal
}
