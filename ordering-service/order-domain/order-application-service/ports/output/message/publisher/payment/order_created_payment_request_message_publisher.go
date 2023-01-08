package payment

import (
	"common-domain/event/publisher"
	"order-domain-core/entity"
	"order-domain-core/event"
)

type OrderCreatedPaymentRequestMessagePublisher interface {
	publisher.DomainEventPublisher[entity.Order, event.OrderCreatedEvent]
}
