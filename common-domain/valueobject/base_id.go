package valueobject

type BaseId[T any] struct {
	Value T
}

func NewBaseId[T any](value T) *BaseId[T] {
	return &BaseId[T]{
		Value: value,
	}
}
