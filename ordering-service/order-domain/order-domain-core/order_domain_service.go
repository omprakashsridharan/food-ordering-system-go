package order_domain_core

import (
	"order-domain-core/entity"
	"order-domain-core/event"
)

type OrderDomainService interface {
	ValidateAndInitiateOrder(order entity.Order, restaurant entity.Restaurant) (event.OrderCreatedEvent, error)
	PayOrder(order entity.Order) (event.OrderPaidEvent, error)
	ApproveOrder(order entity.Order) error
	CancelOrderPayment(order entity.Order, failureMessages []string) (event.OrderCancelledEvent, error)
	CancelOrder(order entity.Order, failureMessages []string) error
}
