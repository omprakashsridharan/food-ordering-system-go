package entity

import (
	"common-domain/entity"
	"common-domain/valueobject"
)

type Customer struct {
	*entity.AggregateRoot[valueobject.CustomerId]
}

func NewCustomer(customerId valueobject.CustomerId) *Customer {
	aggregateRoot := entity.NewAggregateRoot[valueobject.CustomerId](customerId)
	return &Customer{AggregateRoot: aggregateRoot}
}
