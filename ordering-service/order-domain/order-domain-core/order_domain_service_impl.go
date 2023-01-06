package order_domain_core

import (
	"common-domain/exception"
	"order-domain-core/entity"
	"order-domain-core/event"
	"time"
)

type OrderDomainServiceImpl struct {
}

func (o OrderDomainServiceImpl) ValidateAndInitiateOrder(order entity.Order, restaurant entity.Restaurant) (*event.OrderCreatedEvent, error) {
	if !restaurant.Active {
		return nil, exception.OrderDomainException("Inactive restaurant")
	}
	for _, item := range order.Items {
		currentProduct := item.Product
		for _, product := range restaurant.Products {
			if product.Equals(&currentProduct) {
				currentProduct.UpdateWithConfirmedNameAndPrice(product.Name, product.Price)
			}
		}
	}
	err := order.ValidateOrder()
	if err != nil {
		return nil, err
	}
	return event.NewOrderCreatedEvent(&order, time.Now()), nil
}

func (o OrderDomainServiceImpl) PayOrder(order entity.Order) (*event.OrderPaidEvent, error) {
	err := order.Pay()
	if err != nil {
		return nil, err
	}
	return event.NewOrderPaidEvent(&order, time.Now()), nil
}

func (o OrderDomainServiceImpl) ApproveOrder(order entity.Order) error {
	return order.Approve()
}

func (o OrderDomainServiceImpl) CancelOrderPayment(order entity.Order, failureMessages []string) (*event.OrderCancelledEvent, error) {
	err := order.InitCancel(failureMessages)
	if err != nil {
		return nil, err
	}
	return event.NewOrderCancelledEvent(&order, time.Now()), nil
}

func (o OrderDomainServiceImpl) CancelOrder(order entity.Order, failureMessages []string) error {
	return order.Cancel(failureMessages)
}
