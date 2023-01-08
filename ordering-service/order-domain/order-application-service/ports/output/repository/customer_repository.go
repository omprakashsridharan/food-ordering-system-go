package repository

import (
	"github.com/google/uuid"
	"order-domain-core/entity"
)

type CustomerRepository interface {
	FindCustomer(customerId uuid.UUID) (bool, *entity.Customer)
}
