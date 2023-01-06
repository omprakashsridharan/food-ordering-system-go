package entity

type AggregateRoot[ID any] struct {
	*BaseEntity[ID]
}

func NewAggregateRoot[ID any](id ID) *AggregateRoot[ID] {
	baseEntity := NewBaseEntity[ID](id)
	return &AggregateRoot[ID]{BaseEntity: baseEntity}
}
