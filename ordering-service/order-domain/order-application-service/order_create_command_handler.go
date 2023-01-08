package order_application_service

import (
	"order-application-service/dto/create"
	"order-application-service/mapper"
	"order-application-service/ports/output/message/publisher/payment"
)

type OrderCreateCommandHandler struct {
	*OrderCreateHelper
	*mapper.OrderDataMapper
	payment.OrderCreatedPaymentRequestMessagePublisher
}

func NewOrderCreateCommandHandler(orderCreateHelper *OrderCreateHelper, orderDataMapper *mapper.OrderDataMapper, orderCreatedPaymentRequestMessagePublisher payment.OrderCreatedPaymentRequestMessagePublisher) *OrderCreateCommandHandler {
	return &OrderCreateCommandHandler{OrderCreateHelper: orderCreateHelper, OrderDataMapper: orderDataMapper, OrderCreatedPaymentRequestMessagePublisher: orderCreatedPaymentRequestMessagePublisher}
}

func (h *OrderCreateCommandHandler) CreateOrder(command create.CreateOrderCommand) (*create.CreateOrderResponse, error) {
	orderCreatedEvent, err := h.OrderCreateHelper.PersistOrder(command)
	if err != nil {
		return nil, err
	}
	h.OrderCreatedPaymentRequestMessagePublisher.Publish(*orderCreatedEvent)
	return h.OrderDataMapper.OrderToCreateOrderResponse(*orderCreatedEvent.Order, "Order created successfully"), nil
}
