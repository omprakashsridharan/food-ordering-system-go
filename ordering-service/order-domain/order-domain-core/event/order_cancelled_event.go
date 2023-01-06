package event

import "time"

type OrderCancelledEvent struct {
	*OrderEvent
}

func NewOrderCancelledEvent(createdAt time.Time) *OrderCancelledEvent {
	orderEvent := NewOrderEvent(createdAt)
	return &OrderCancelledEvent{OrderEvent: orderEvent}
}
