package valueobject

import "github.com/google/uuid"

type ProductId struct {
	*BaseId[uuid.UUID]
}

func NewProductId(value uuid.UUID) *ProductId {
	baseId := NewBaseId[uuid.UUID](value)
	return &ProductId{BaseId: baseId}
}
