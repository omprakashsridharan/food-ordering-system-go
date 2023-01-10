package track

import "github.com/google/uuid"

type TrackOrderQuery struct {
	OrderTrackingId uuid.UUID
}

func NewTrackOrderQuery(orderTrackingId uuid.UUID) *TrackOrderQuery {
	return &TrackOrderQuery{OrderTrackingId: orderTrackingId}
}
