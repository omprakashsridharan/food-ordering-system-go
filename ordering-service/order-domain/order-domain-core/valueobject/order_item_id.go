package valueobject

import (
	"common-domain/valueobject"
)

type OrderItemId struct {
	valueobject.BaseId[int64]
}

func NewOrderItemId(value int64) *OrderItemId {
	baseId := valueobject.BaseId[int64]{
		Value: value,
	}
	return &OrderItemId{BaseId: baseId}
}
