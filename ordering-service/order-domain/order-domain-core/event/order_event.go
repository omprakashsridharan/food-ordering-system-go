package event

import (
	"order-domain-core/entity"
	"time"
)

type OrderEvent struct {
	*entity.Order
	CreatedAt time.Time
}

func (o OrderEvent) GetCreatedAt() time.Time {
	return o.CreatedAt
}

func NewOrderEvent(order *entity.Order, createdAt time.Time) *OrderEvent {
	return &OrderEvent{Order: order, CreatedAt: createdAt}
}
