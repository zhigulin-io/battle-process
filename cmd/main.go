package main

import (
	"battle-process/internal/game"
)

func main() {
	unitA := game.NewUnit("Test Unit A", 4, 4)
	playerA := game.NewPlayer()
	playerA.AddUnit(unitA)

	unitB := game.NewUnit("Test Unit B", 4, 4)
	playerB := game.NewPlayer()
	playerB.AddUnit(unitB)

	g := game.Game{
		ActivePlayer:  playerA,
		PassivePlayer: playerB,
	}
	actionChan, responseChan := g.Run()

	*actionChan <- game.Action{
		PlayerID:       playerA.ID,
		ActivateAction: &game.ActivateAction{UnitID: unitA.GetID()},
	}

	_ = <-*responseChan

	*actionChan <- game.Action{
		PlayerID:       playerB.ID,
		ActivateAction: &game.ActivateAction{UnitID: unitB.GetID()},
	}

	_ = <-*responseChan
	_ = <-*responseChan
	//fmt.Println("Success:", response.Success)
	//if response.Message != nil {
	//	fmt.Println("Message:", *response.Message)
	//}
}
