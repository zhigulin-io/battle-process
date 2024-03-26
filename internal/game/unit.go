package game

import "github.com/google/uuid"

type Unit struct {
	id      uuid.UUID
	name    string
	defence int
	quality int
	wounds  int
}

func NewUnit(name string, defence, quality int) *Unit {
	return &Unit{
		id:      uuid.New(),
		name:    name,
		defence: defence,
		quality: quality,
	}
}

func (u *Unit) GetID() uuid.UUID {
	return u.id
}
