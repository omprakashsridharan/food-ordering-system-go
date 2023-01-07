package track

import (
	"common-domain/valueobject"
	"github.com/google/uuid"
)

type TrackOrderResponse struct {
	OrderTrackingId uuid.UUID
	OrderStatus     valueobject.OrderStatus
	FailureMessages []string
}

func NewTrackOrderResponse(orderTrackingId uuid.UUID, orderStatus valueobject.OrderStatus, failureMessages []string) *TrackOrderResponse {
	return &TrackOrderResponse{OrderTrackingId: orderTrackingId, OrderStatus: orderStatus, FailureMessages: failureMessages}
}
