package event

import (
	"order-domain-core/entity"
	"time"
)

type OrderPaidEvent struct {
	*OrderEvent
}

func NewOrderPaidEvent(order *entity.Order, createdAt time.Time) *OrderPaidEvent {
	orderEvent := NewOrderEvent(order, createdAt)
	return &OrderPaidEvent{OrderEvent: orderEvent}
}
