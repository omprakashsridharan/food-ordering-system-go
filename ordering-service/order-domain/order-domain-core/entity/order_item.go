package entity

import (
	"common-domain/entity"
	commonValueObject "common-domain/valueobject"
	orderDomainValueObject "order-domain-core/valueobject"
)

type OrderItem struct {
	*entity.BaseEntity[orderDomainValueObject.OrderItemId]
	OrderId  commonValueObject.OrderId
	Product  Product
	Quantity int64
	Price    commonValueObject.Money
	SubTotal commonValueObject.Money
}

func NewOrderItem(orderItemId orderDomainValueObject.OrderItemId, orderId commonValueObject.OrderId, product Product, quantity int64, price commonValueObject.Money, subTotal commonValueObject.Money) *OrderItem {
	baseEntity := entity.NewBaseEntity[orderDomainValueObject.OrderItemId](orderItemId)
	return &OrderItem{BaseEntity: baseEntity, OrderId: orderId, Product: product, Quantity: quantity, Price: price, SubTotal: subTotal}
}

func (oi *OrderItem) PriceValid() bool {
	return oi.Price.GreaterThanZero() && oi.Price == oi.Product.Price && oi.Price.Multiply(oi.Quantity).Equals(&oi.SubTotal)
}
