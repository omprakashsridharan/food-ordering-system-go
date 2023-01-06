package publisher

import "common-domain/event"

type DomainEventPublisher[E any, T event.DomainEvent[E]] interface {
	Publish(domainEvent T)
}
