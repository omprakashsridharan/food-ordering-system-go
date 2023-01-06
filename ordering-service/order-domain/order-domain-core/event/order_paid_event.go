package event

import "time"

type OrderPaidEvent struct {
	*OrderEvent
}

func NewOrderPaidEvent(createdAt time.Time) *OrderPaidEvent {
	orderEvent := NewOrderEvent(createdAt)
	return &OrderPaidEvent{OrderEvent: orderEvent}
}
