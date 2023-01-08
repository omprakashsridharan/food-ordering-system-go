package service

import (
	"order-application-service/dto/create"
	"order-application-service/dto/track"
)

type OrderApplicationService interface {
	CreateOrder(command create.CreateOrderCommand) *create.CreateOrderResponse
	TrackOrder(query track.TrackOrderQuery) *track.TrackOrderResponse
}
