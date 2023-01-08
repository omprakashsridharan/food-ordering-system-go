package payment

import (
	"common-domain/event/publisher"
	"order-domain-core/entity"
	"order-domain-core/event"
)

type OrderCancelledPaymentRequestMessagePublisher interface {
	publisher.DomainEventPublisher[entity.Order, event.OrderCancelledEvent]
}
