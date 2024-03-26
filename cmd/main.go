package main

import (
	"battle-process/internal/game"
	"github.com/google/uuid"
)

func main() {
	playerA := game.Player{
		ID:             uuid.New(),
		Score:          0,
		ActiveUnit:     nil,
		AwaitingUnits:  make(map[uuid.UUID]*game.Unit),
		ActivatedUnits: make(map[uuid.UUID]*game.Unit),
	}
	playerB := game.Player{
		ID:             uuid.New(),
		Score:          0,
		ActiveUnit:     nil,
		AwaitingUnits:  nil,
		ActivatedUnits: nil,
	}

	unit := game.Unit{
		ID:      uuid.New(),
		Name:    "Test Unit",
		Defence: 4,
		Quality: 4,
		State:   "test state",
		Wounds:  0,
	}
	playerA.AwaitingUnits[unit.ID] = &unit

	g := game.Game{
		ActivePlayer:  &playerA,
		PassivePlayer: &playerB,
	}
	actionChan, responseChan := g.Run()

	*actionChan <- game.Action{
		PlayerID:       playerA.ID,
		ActivateAction: &game.ActivateAction{UnitID: unit.ID},
	}

	_ = <-*responseChan
	//fmt.Println("Success:", response.Success)
	//if response.Message != nil {
	//	fmt.Println("Message:", *response.Message)
	//}
}
