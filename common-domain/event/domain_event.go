package event

import "time"

type DomainEvent[T any] struct {
	CreatedAt time.Time
}

func NewDomainEvent[T any](createdAt time.Time) *DomainEvent[T] {
	return &DomainEvent[T]{CreatedAt: createdAt}
}
