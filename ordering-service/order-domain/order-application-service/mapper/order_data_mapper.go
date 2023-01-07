package mapper

import (
	"common-domain/valueobject"
	"github.com/google/uuid"
	"order-application-service/dto/create"
	"order-application-service/dto/track"
	"order-domain-core/entity"
	orderValueObject "order-domain-core/valueobject"
)

type OrderDataMapper struct {
}

func (odm *OrderDataMapper) CreateOrderCommandToRestaurant(command create.CreateOrderCommand) *entity.Restaurant {
	var products = make([]entity.Product, 0)
	for _, item := range command.Items {
		product := entity.NewProduct(*valueobject.NewProductId(item.ProductId), "", *valueobject.ZERO)
		products = append(products, *product)
	}
	restaurantId := valueobject.NewRestaurantId(command.RestaurantId)
	return entity.NewRestaurant(*restaurantId, products, true)
}

func (odm *OrderDataMapper) CreateOrderCommandToOrder(command create.CreateOrderCommand) *entity.Order {
	customerId := valueobject.NewCustomerId(command.CustomerId)
	restaurantId := valueobject.NewRestaurantId(command.RestaurantId)
	deliveryAddress := orderValueObject.NewStreetAddress(uuid.New(), command.Address.Street, command.Address.PostalCode, command.Address.City)
	price := valueobject.NewMoneyFromDecimal(command.Price)
	var items []entity.OrderItem
	orderId := valueobject.NewOrderId(uuid.New())
	trackingId := orderValueObject.NewTrackingId(uuid.New())
	for index, item := range command.Items {
		product := entity.NewProduct(*valueobject.NewProductId(item.ProductId), "", *valueobject.ZERO)
		price := valueobject.NewMoneyFromDecimal(item.Price)
		quantity := item.Quantity
		subTotal := valueobject.NewMoneyFromDecimal(item.SubTotal)
		orderItemId := orderValueObject.NewOrderItemId(int64(index))
		orderItem := entity.NewOrderItem(*orderItemId, *orderId, *product, int64(quantity), *price, *subTotal)
		items = append(items, *orderItem)
	}
	order := entity.NewOrder(*orderId, *customerId, *restaurantId, *deliveryAddress, *price, items, *trackingId, valueobject.ORDER_STATUS_PENDING, []string{})
	return order
}

func (odm *OrderDataMapper) OrderToCreateOrderResponse(order entity.Order, message string) *create.CreateOrderResponse {
	return create.NewCreateOrderResponse(order.TrackingId.Value, order.OrderStatus, message)
}

func (odm *OrderDataMapper) OrderToTrackOrderResponse(order entity.Order) *track.TrackOrderResponse {
	return track.NewTrackOrderResponse(order.TrackingId.Value, order.OrderStatus, order.FailureMessages)
}
