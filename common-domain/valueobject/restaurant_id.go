package valueobject

import "github.com/google/uuid"

type RestaurantId struct {
	BaseId[uuid.UUID]
}

func NewRestaurantId(value uuid.UUID) *RestaurantId {
	baseId := BaseId[uuid.UUID]{value}
	return &RestaurantId{BaseId: baseId}
}
