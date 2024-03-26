package game

import (
	"github.com/google/uuid"
	"log"
)

type ActivateAction struct {
	UnitID uuid.UUID
}

type HoldAction struct {
}

type AdvanceAction struct {
}

type RushAction struct {
}

type ChargeAction struct {
}

type ShootAction struct {
}

type FightAction struct {
}

type Action struct {
	PlayerID       uuid.UUID
	ActivateAction *ActivateAction
	HoldAction     *HoldAction
	AdvanceAction  *AdvanceAction
	RushAction     *RushAction
	ChargeAction   *ChargeAction
	ShootAction    *ShootAction
	FightAction    *FightAction
}

type Response struct {
	Success bool
	Message *string
}

type Game struct {
	ActivePlayer  *Player
	PassivePlayer *Player
}

func (g *Game) Run() (*chan Action, *chan Response) {
	actionChan := make(chan Action)
	responseChan := make(chan Response)

	go func() {
		g.gameProcess(&actionChan, &responseChan)
	}()

	return &actionChan, &responseChan
}

func (g *Game) gameProcess(actionChan *chan Action, responseChan *chan Response) {
	// Preparation
	// Deployment Roll-off

	for i := 0; i < 4; i++ {
		// 1. Round Begin
		// Caster points
		log.Println("===== Round Begin =====")
		// 2. Turns
		for len(g.ActivePlayer.AwaitingUnits) != 0 || len(g.PassivePlayer.AwaitingUnits) != 0 {
			log.Println("===== Turn Begin ======")
			log.Printf("Turn of Player %s", g.ActivePlayer.ID)
			log.Printf("Player %s has %d awaiting unit(s)", g.ActivePlayer.ID, len(g.ActivePlayer.AwaitingUnits))
			if len(g.ActivePlayer.AwaitingUnits) != 0 {
				g.activatePhase(actionChan, responseChan)

				//performedAction := g.movePhase(actionChan)
				//
				//if performedAction == "hold" || performedAction == "advance" {
				//	g.shootingPhase(actionChan)
				//} else if performedAction == "charge" {
				//	g.shootingPhase(actionChan)
				//}

				g.ActivePlayer.ActivatedUnits[g.ActivePlayer.ActiveUnit.id] = g.ActivePlayer.ActiveUnit
				g.ActivePlayer.ActiveUnit = nil
			}
			log.Println("===== End of Turn =====")
			g.switchPlayers()
		}

		// 3. Round End
		// Scoring
		g.ActivePlayer.AwaitingUnits = g.ActivePlayer.ActivatedUnits
		g.ActivePlayer.ActivatedUnits = make(map[uuid.UUID]*Unit, len(g.ActivePlayer.AwaitingUnits))

		g.PassivePlayer.AwaitingUnits = g.PassivePlayer.ActivatedUnits
		g.PassivePlayer.ActivatedUnits = make(map[uuid.UUID]*Unit, len(g.PassivePlayer.AwaitingUnits))
		log.Println("===== Round End =====")
	}
}

func (g *Game) switchPlayers() {
	tmp := g.ActivePlayer
	g.ActivePlayer = g.PassivePlayer
	g.PassivePlayer = tmp
}

func (g *Game) activatePhase(actionChan *chan Action, responseChan *chan Response) {
	for {
		action := <-*actionChan
		if action.PlayerID != g.ActivePlayer.ID {
			message := "invalid player id"
			*responseChan <- Response{Message: &message}
			continue
		}

		if action.ActivateAction == nil {
			message := "activation action required"
			*responseChan <- Response{Message: &message}
			continue
		}

		unit, ok := g.ActivePlayer.AwaitingUnits[action.ActivateAction.UnitID]
		if !ok {
			message := "invalid unit id"
			*responseChan <- Response{Message: &message}
			continue
		}

		g.ActivePlayer.ActiveUnit = unit
		delete(g.ActivePlayer.AwaitingUnits, action.ActivateAction.UnitID)
		*responseChan <- Response{Success: true}
		return
	}
}

func (g *Game) movePhase(actionChan *chan Action) string {
	for {
		action := <-*actionChan
		if action.PlayerID != g.ActivePlayer.ID {
			// write error response
			continue
		}

		if action.HoldAction != nil {
			// process hold action
			return "hold"
		}

		if action.AdvanceAction != nil {
			// process advance action
			return "advance"
		}

		if action.RushAction != nil {
			// process rush action
			return "rush"
		}

		if action.ChargeAction != nil {
			// process charge action
			return "charge"
		}

		// write error response
	}
}

func (g *Game) shootingPhase(actionChan *chan Action) {
	for {
		action := <-*actionChan
		if action.PlayerID != g.ActivePlayer.ID {
			// write error response
			continue
		}

		if action.ShootAction == nil {
			// process hold action
			continue
		}

		// process shooting action
		// 1. determine attacks
		// 2. roll to hit
		// 3. roll to block
		// 4. check wound effects
		return
	}
}

func (g *Game) fightingPhase(actionChan *chan Action) {
	for {
		action := <-*actionChan
		if action.PlayerID != g.ActivePlayer.ID {
			// write error response
			continue
		}

		if action.FightAction == nil {
			// process hold action
			continue
		}

		// process shooting action
		// 1. determine attacks
		// 2. roll to hit
		// 3. roll to block
		// 4. check wound effects
		return
	}
}
