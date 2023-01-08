package repository

import (
	"order-domain-core/entity"
	"order-domain-core/valueobject"
)

type OrderRepository interface {
	Save(order entity.Order) (bool, *entity.Order)
	FindByTrackingId(id valueobject.TrackingId) (bool, *entity.Order)
}
