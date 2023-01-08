package order_application_service

import (
	"common-domain/exception"
	"github.com/google/uuid"
	"order-application-service/dto/create"
	"order-application-service/mapper"
	"order-application-service/ports/output/repository"
	orderDomainCore "order-domain-core"
	"order-domain-core/entity"
	"order-domain-core/event"
)

type OrderCreateHelper struct {
	orderDomainCore.OrderDomainService
	repository.OrderRepository
	repository.CustomerRepository
	repository.RestaurantRepository
	*mapper.OrderDataMapper
}

func NewOrderCreateHelper(orderDomainService orderDomainCore.OrderDomainService, orderRepository repository.OrderRepository, customerRepository repository.CustomerRepository, restaurantRepository repository.RestaurantRepository, orderDataMapper *mapper.OrderDataMapper) *OrderCreateHelper {
	return &OrderCreateHelper{OrderDomainService: orderDomainService, OrderRepository: orderRepository, CustomerRepository: customerRepository, RestaurantRepository: restaurantRepository, OrderDataMapper: orderDataMapper}
}

func (och *OrderCreateHelper) PersistOrder(command create.CreateOrderCommand) (*event.OrderCreatedEvent, error) {
	err := och.CheckCustomer(command.CustomerId)
	if err != nil {
		return nil, err
	}
	restaurant, err := och.CheckRestaurant(command)
	if err != nil {
		return nil, err
	}
	order := och.CreateOrderCommandToOrder(command)
	orderCreatedEvent, err := och.OrderDomainService.ValidateAndInitiateOrder(*order, *restaurant)
	if err != nil {
		return nil, err
	}
	return orderCreatedEvent, nil
}

func (och *OrderCreateHelper) SaveOrder(order entity.Order) (*entity.Order, error) {
	ok, o := och.OrderRepository.Save(order)
	if !ok {
		return nil, exception.OrderDomainException("Could not save order")
	}
	return o, nil
}

func (och *OrderCreateHelper) CheckRestaurant(command create.CreateOrderCommand) (*entity.Restaurant, error) {
	restaurant := och.OrderDataMapper.CreateOrderCommandToRestaurant(command)
	ok, r := och.RestaurantRepository.FindRestaurantInfo(*restaurant)
	if !ok {
		return nil, exception.OrderDomainException("Could not find restaurant with restaurant id")
	}
	return r, nil
}

func (och *OrderCreateHelper) CheckCustomer(customerId uuid.UUID) error {
	ok, _ := och.CustomerRepository.FindCustomer(customerId)
	if !ok {
		return exception.OrderDomainException("Could not find customer with customer id")
	}
	return nil
}
