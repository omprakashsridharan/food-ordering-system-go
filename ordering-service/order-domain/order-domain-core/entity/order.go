package entity

import (
	"common-domain/entity"
	"common-domain/exception"
	commonValueObject "common-domain/valueobject"
	orderDomainValueObject "order-domain-core/valueobject"
)

type Order struct {
	*entity.AggregateRoot[commonValueObject.OrderId]
	CustomerId      commonValueObject.CustomerId
	RestaurantId    commonValueObject.RestaurantId
	StreetAddress   orderDomainValueObject.StreetAddress
	Price           commonValueObject.Money
	Items           []OrderItem
	TrackingId      orderDomainValueObject.TrackingId
	OrderStatus     commonValueObject.OrderStatus
	FailureMessages []string
}

func NewOrder(orderId commonValueObject.OrderId, customerId commonValueObject.CustomerId, restaurantId commonValueObject.RestaurantId, streetAddress orderDomainValueObject.StreetAddress, price commonValueObject.Money, items []OrderItem, trackingId orderDomainValueObject.TrackingId, orderStatus commonValueObject.OrderStatus, failureMessages []string) *Order {
	aggregateRoot := entity.NewAggregateRoot[commonValueObject.OrderId](orderId)
	return &Order{AggregateRoot: aggregateRoot, CustomerId: customerId, RestaurantId: restaurantId, StreetAddress: streetAddress, Price: price, Items: items, TrackingId: trackingId, OrderStatus: orderStatus, FailureMessages: failureMessages}
}

func (o *Order) ValidateOrder() error {
	err := o.ValidateTotalPrice()
	if err != nil {
		return err
	}

	err = o.ValidateItemsPrice()
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) Pay() error {
	if o.OrderStatus != commonValueObject.ORDER_STATUS_PENDING {
		return exception.OrderDomainException("Order not in correct state for pay operation")
	}
	o.OrderStatus = commonValueObject.ORDER_STATUS_PAID
	return nil
}

func (o *Order) Approve() error {
	if o.OrderStatus != commonValueObject.ORDER_STATUS_PAID {
		return exception.OrderDomainException("Order not in correct state for approve operation")
	}
	o.OrderStatus = commonValueObject.ORDER_STATUS_APPROVED
	return nil
}

func (o *Order) InitCancel(failureMessages []string) error {
	if o.OrderStatus != commonValueObject.ORDER_STATUS_PAID {
		return exception.OrderDomainException("Order not in correct state for init cancel operation")
	}
	o.OrderStatus = commonValueObject.ORDER_STATUS_CANCELLING
	o.updateFailureMessages(failureMessages)
	return nil
}

func (o *Order) Cancel(failureMessages []string) error {
	if !(o.OrderStatus == commonValueObject.ORDER_STATUS_CANCELLING || o.OrderStatus == commonValueObject.ORDER_STATUS_PENDING) {
		return exception.OrderDomainException("Order not in correct state for cancel operation")
	}
	o.OrderStatus = commonValueObject.ORDER_STATUS_CANCELLED
	o.updateFailureMessages(failureMessages)
	return nil
}

func (o *Order) updateFailureMessages(messages []string) {
	if o.FailureMessages == nil {
		o.FailureMessages = make([]string, len(messages))
	}
	o.FailureMessages = append(o.FailureMessages, messages...)
}

func (o *Order) ValidateItemsPrice() error {
	var orderItemsTotalPrice = commonValueObject.NewMoney(0.0)
	for _, item := range o.Items {
		isItemPriceValid := item.PriceValid()
		if !isItemPriceValid {
			return exception.OrderDomainException("Order item price invalid")
		} else {
			orderItemsTotalPrice = orderItemsTotalPrice.Add(&item.SubTotal)
		}
	}
	if !o.Price.Equals(orderItemsTotalPrice) {
		return exception.OrderDomainException("Total does not match the individual items total")
	}
	return nil
}

func (o *Order) ValidateTotalPrice() error {
	if !o.Price.GreaterThanZero() {
		return exception.OrderDomainException("Total price must be greater than zero")
	}
	return nil
}
