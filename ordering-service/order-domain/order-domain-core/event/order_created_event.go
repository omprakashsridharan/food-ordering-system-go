package event

import "time"

type OrderCreatedEvent struct {
	*OrderEvent
}

func NewOrderCreatedEvent(createdAt time.Time) *OrderCreatedEvent {
	orderEvent := NewOrderEvent(createdAt)
	return &OrderCreatedEvent{OrderEvent: orderEvent}
}
