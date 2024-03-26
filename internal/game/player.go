package game

import "github.com/google/uuid"

type Player struct {
	ID uuid.UUID

	Score int

	ActiveUnit     *Unit
	AwaitingUnits  map[uuid.UUID]*Unit
	ActivatedUnits map[uuid.UUID]*Unit
}

func NewPlayer() *Player {
	return &Player{
		ID:             uuid.New(),
		AwaitingUnits:  make(map[uuid.UUID]*Unit),
		ActivatedUnits: make(map[uuid.UUID]*Unit),
	}
}

func (p *Player) AddUnit(unit *Unit) {
	p.AwaitingUnits[unit.id] = unit
}
