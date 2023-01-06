package event

import (
	"common-domain/event"
	"order-domain-core/entity"
	"time"
)

type OrderEvent struct {
	*event.DomainEvent[entity.Order]
	*entity.Order
}

func NewOrderEvent(order *entity.Order, createdAt time.Time) *OrderEvent {
	domainEvent := event.NewDomainEvent[entity.Order](createdAt)
	return &OrderEvent{DomainEvent: domainEvent, Order: order}
}
