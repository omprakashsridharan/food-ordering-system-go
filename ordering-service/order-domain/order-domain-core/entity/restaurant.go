package entity

import (
	"common-domain/entity"
	commonValueObject "common-domain/valueobject"
)

type Restaurant struct {
	*entity.AggregateRoot[commonValueObject.RestaurantId]
	Products []Product
	Active   bool
}

func NewRestaurant(restaurantId commonValueObject.RestaurantId, products []Product, active bool) *Restaurant {
	aggregateRoot := entity.NewAggregateRoot[commonValueObject.RestaurantId](restaurantId)
	return &Restaurant{AggregateRoot: aggregateRoot, Products: products, Active: active}
}
