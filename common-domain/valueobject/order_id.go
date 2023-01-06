package valueobject

import "github.com/google/uuid"

type OrderId struct {
	*BaseId[uuid.UUID]
}

func NewOrderId(value uuid.UUID) *OrderId {
	baseId := NewBaseId[uuid.UUID](value)
	return &OrderId{BaseId: baseId}
}
