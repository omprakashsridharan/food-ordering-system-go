package valueobject

import (
	"common-domain/valueobject"
	"github.com/google/uuid"
)

type TrackingId struct {
	valueobject.BaseId[uuid.UUID]
}

func NewTrackingId(value uuid.UUID) *TrackingId {
	baseId := valueobject.BaseId[uuid.UUID]{value}
	return &TrackingId{BaseId: baseId}
}
