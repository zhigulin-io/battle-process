package game

import (
	"errors"
)

type UnitState string

const (
	UnitStateAwaiting  UnitState = "awaiting"
	UnitStateActive    UnitState = "active"
	UnitStateActivated UnitState = "activated"
)

type Unit struct {
	Owner int       `json:"owner"`
	State UnitState `json:"state"`
}

type Game struct {
	Round        int    `json:"round"`
	ActivePlayer int    `json:"activePlayer"`
	UnitList     []Unit `json:"unitList"`
}

func NewGame() *Game {
	g := Game{
		Round:        0,
		ActivePlayer: 0,
		UnitList:     make([]Unit, 4),
	}

	for i := range g.UnitList {
		g.UnitList[i] = Unit{
			Owner: i % 2,
			State: UnitStateAwaiting,
		}
	}

	return &g
}

func (g *Game) GetState() any {
	return g
}

func (g *Game) StartActivation(playerID int, unitID int) error {
	if unitID < 0 || unitID >= len(g.UnitList) {
		return errors.New("unit id is too small or too large")
	}

	if playerID != g.ActivePlayer {
		return errors.New("the player is not active now")
	}

	if g.UnitList[unitID].Owner != playerID {
		return errors.New("other player owns this unit")
	}

	for i := range g.UnitList {
		if g.UnitList[i].Owner == playerID && g.UnitList[i].State == UnitStateActive {
			return errors.New("the player has active unit already")
		}
	}

	if g.UnitList[unitID].State == UnitStateActivated {
		return errors.New("this unit is activated already")
	}

	g.UnitList[unitID].State = UnitStateActive

	return nil
}

func (g *Game) StopActivation(playerID int) error {
	remainingUnits := 0
	unitID := -1
	for i := range g.UnitList {
		if g.UnitList[i].State == UnitStateAwaiting {
			remainingUnits++
		}
		if g.UnitList[i].Owner == playerID && g.UnitList[i].State == UnitStateActive {
			unitID = i
		}
	}
	if unitID == -1 {
		return errors.New("the player hasn't active unit")
	}

	g.UnitList[unitID].State = UnitStateActivated
	g.ActivePlayer = (g.ActivePlayer + 1) % 2
	if remainingUnits == 0 {
		g.Round++
		for i := range g.UnitList {
			g.UnitList[i].State = UnitStateAwaiting
		}
	}

	return nil
}
