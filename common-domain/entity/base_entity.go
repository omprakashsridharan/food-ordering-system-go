package entity

type BaseEntity[ID any] struct {
	Id ID
}

func NewBaseEntity[ID any](id ID) *BaseEntity[ID] {
	return &BaseEntity[ID]{Id: id}
}
