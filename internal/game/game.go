package game

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Phase string

const (
	PhaseActivation = "activation"
	PhaseShooting   = "shooting"
	PhaseFighting   = "fighting"
	PhaseMoving     = "moving"
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

	phase Phase
}

func (g *Game) Run() (chan Action, chan Response) {
	actionChan := make(chan Action)
	responseChan := make(chan Response)

	go func() {
		g.gameProcess(actionChan, responseChan)
	}()

	return actionChan, responseChan
}

func (g *Game) gameProcess(actionChan chan Action, responseChan chan Response) {
	for i := 0; i < 4; i++ {
		g.roundBegin()

		for len(g.ActivePlayer.AwaitingUnits) != 0 || len(g.PassivePlayer.AwaitingUnits) != 0 {
			if len(g.ActivePlayer.AwaitingUnits) != 0 {
				g.turnController(actionChan, responseChan)
			}
			g.switchPlayers()
		}

		g.roundEnd()
	}
}

func (g *Game) roundBegin() {
	g.phase = PhaseActivation
}

func (g *Game) roundEnd() {
	g.ActivePlayer.AwaitingUnits = g.ActivePlayer.ActivatedUnits
	g.ActivePlayer.ActivatedUnits = make(map[uuid.UUID]*Unit, len(g.ActivePlayer.AwaitingUnits))

	g.PassivePlayer.AwaitingUnits = g.PassivePlayer.ActivatedUnits
	g.PassivePlayer.ActivatedUnits = make(map[uuid.UUID]*Unit, len(g.PassivePlayer.AwaitingUnits))
}

func (g *Game) switchPlayers() {
	tmp := g.ActivePlayer
	g.ActivePlayer = g.PassivePlayer
	g.PassivePlayer = tmp
}

func (g *Game) turnController(actionChan chan Action, responseChan chan Response) {
	for {
		action := <-actionChan
		if action.PlayerID != g.ActivePlayer.ID {
			message := "invalid player id"
			responseChan <- Response{Message: &message}
			continue
		}

		var turnEnd bool
		switch g.phase {
		case PhaseActivation:
			err := g.activatePhaseHandler(action)
			if err != nil {
				message := err.Error()
				responseChan <- Response{Message: &message}
				continue
			}
			g.phase = PhaseMoving
		case PhaseMoving:
		case PhaseShooting:
			turnEnd = true
		case PhaseFighting:
			turnEnd = true
		default:
			panic(fmt.Sprintf("invalid phase: %s", g.phase))
		}

		if turnEnd {
			break
		}
	}

	g.ActivePlayer.ActivatedUnits[g.ActivePlayer.ActiveUnit.id] = g.ActivePlayer.ActiveUnit
	g.ActivePlayer.ActiveUnit = nil
}

func (g *Game) activatePhaseHandler(action Action) error {
	if action.ActivateAction == nil {
		return errors.New("activation action required")
	}

	unit, ok := g.ActivePlayer.AwaitingUnits[action.ActivateAction.UnitID]
	if !ok {
		return errors.New("invalid unit id")
	}

	g.ActivePlayer.ActiveUnit = unit
	delete(g.ActivePlayer.AwaitingUnits, action.ActivateAction.UnitID)
	return nil
}

func (g *Game) movePhase(actionChan chan Action) string {
	for {
		action := <-actionChan

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

func (g *Game) shootingPhase(actionChan chan Action) {
	for {
		action := <-actionChan

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

func (g *Game) fightingPhase(actionChan chan Action) {
	for {
		action := <-actionChan

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
