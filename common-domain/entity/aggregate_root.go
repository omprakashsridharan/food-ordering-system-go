package entity

type AggregateRoot[ID any] struct {
	BaseEntity[ID]
}
