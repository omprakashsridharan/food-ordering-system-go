package valueobject

import "github.com/google/uuid"

type CustomerId struct {
	BaseId[uuid.UUID]
}

func NewCustomerId(value uuid.UUID) *CustomerId {
	baseId := BaseId[uuid.UUID]{value}
	return &CustomerId{BaseId: baseId}
}
