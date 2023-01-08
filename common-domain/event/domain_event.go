package event

import "time"

type DomainEvent[T any] interface {
	GetCreatedAt() time.Time
}
