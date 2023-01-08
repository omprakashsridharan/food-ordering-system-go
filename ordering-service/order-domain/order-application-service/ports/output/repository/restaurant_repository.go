package repository

import "order-domain-core/entity"

type RestaurantRepository interface {
	FindRestaurantInfo(restaurant entity.Restaurant) (bool, *entity.Restaurant)
}
