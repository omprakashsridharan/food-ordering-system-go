package order_application_service

import (
	"common-domain/exception"
	"order-application-service/dto/track"
	"order-application-service/mapper"
	"order-application-service/ports/output/repository"
	"order-domain-core/valueobject"
)

type OrderTrackCommandHandler struct {
	repository.OrderRepository
	*mapper.OrderDataMapper
}

func NewOrderTrackCommandHandler(orderRepository repository.OrderRepository, orderDataMapper *mapper.OrderDataMapper) *OrderTrackCommandHandler {
	return &OrderTrackCommandHandler{OrderRepository: orderRepository, OrderDataMapper: orderDataMapper}
}

func (h *OrderTrackCommandHandler) TrackOrder(query track.TrackOrderQuery) (*track.TrackOrderResponse, error) {
	ok, order := h.OrderRepository.FindByTrackingId(*valueobject.NewTrackingId(query.OrderTrackingId))
	if !ok {
		return nil, exception.OrderNotFoundException("Order not found")
	}
	return h.OrderDataMapper.OrderToTrackOrderResponse(*order), nil
}
