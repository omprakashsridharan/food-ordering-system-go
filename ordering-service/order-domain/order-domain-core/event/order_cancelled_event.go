package event

import (
	"order-domain-core/entity"
	"time"
)

type OrderCancelledEvent struct {
	*OrderEvent
}

func NewOrderCancelledEvent(order *entity.Order, createdAt time.Time) *OrderCancelledEvent {
	orderEvent := NewOrderEvent(order, createdAt)
	return &OrderCancelledEvent{OrderEvent: orderEvent}
}
