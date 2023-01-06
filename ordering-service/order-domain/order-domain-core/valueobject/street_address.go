package valueobject

import "github.com/google/uuid"

type StreetAddress struct {
	Id         uuid.UUID
	Street     string
	PostalCode string
	City       string
}

func NewStreetAddress(id uuid.UUID, street string, postalCode string, city string) *StreetAddress {
	return &StreetAddress{Id: id, Street: street, PostalCode: postalCode, City: city}
}
