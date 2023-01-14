package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderAddressEntity struct {
	gorm.Model
	Id         uuid.UUID
	Street     string
	PostalCode string
	City       string
	OrderId    string
}
