package entity

import (
	"common-domain/entity"
	commonValueObject "common-domain/valueobject"
)

type Product struct {
	*entity.BaseEntity[commonValueObject.ProductId]
	Name  string
	Price commonValueObject.Money
}

func NewProduct(productId commonValueObject.ProductId, name string, price commonValueObject.Money) *Product {
	baseEntity := entity.NewBaseEntity[commonValueObject.ProductId](productId)
	return &Product{BaseEntity: baseEntity, Name: name, Price: price}
}
