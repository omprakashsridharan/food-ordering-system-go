package event

import (
	"order-domain-core/entity"
	"time"
)

type OrderCreatedEvent struct {
	*OrderEvent
}

func NewOrderCreatedEvent(order *entity.Order, createdAt time.Time) *OrderCreatedEvent {
	orderEvent := NewOrderEvent(order, createdAt)
	return &OrderCreatedEvent{OrderEvent: orderEvent}
}
