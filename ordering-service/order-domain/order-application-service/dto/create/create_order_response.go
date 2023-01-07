package create

import (
	"common-domain/valueobject"
	"github.com/google/uuid"
)

type CreateOrderResponse struct {
	OrderTrackingId uuid.UUID
	OrderStatus     valueobject.OrderStatus
	Message         string
}

func NewCreateOrderResponse(orderTrackingId uuid.UUID, orderStatus valueobject.OrderStatus, message string) *CreateOrderResponse {
	return &CreateOrderResponse{OrderTrackingId: orderTrackingId, OrderStatus: orderStatus, Message: message}
}
