package order_application_service

import (
	"order-application-service/dto/create"
	"order-application-service/dto/track"
)

type OrderApplicationServiceImpl struct {
	*OrderCreateCommandHandler
	*OrderTrackCommandHandler
}

func (o *OrderApplicationServiceImpl) CreateOrder(command create.CreateOrderCommand) (*create.CreateOrderResponse, error) {
	return o.OrderCreateCommandHandler.CreateOrder(command)
}

func (o *OrderApplicationServiceImpl) TrackOrder(query track.TrackOrderQuery) (*track.TrackOrderResponse, error) {
	return o.OrderTrackCommandHandler.TrackOrder(query)
}

func NewOrderApplicationServiceImpl(orderCreateCommandHandler *OrderCreateCommandHandler, orderTrackCommandHandler *OrderTrackCommandHandler) *OrderApplicationServiceImpl {
	return &OrderApplicationServiceImpl{OrderCreateCommandHandler: orderCreateCommandHandler, OrderTrackCommandHandler: orderTrackCommandHandler}
}
