package restaurantapproval

import (
	"common-domain/event/publisher"
	"order-domain-core/entity"
	"order-domain-core/event"
)

type OrderPaidRestaurantRequestMessagePublisher interface {
	publisher.DomainEventPublisher[entity.Order, event.OrderPaidEvent]
}
