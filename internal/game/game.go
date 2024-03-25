package game

import "github.com/google/uuid"

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

type Unit struct {
	Name    string
	Defence int
	Quality int
	State   string
	Wounds  int
}

type Player struct {
	ID uuid.UUID

	Score int

	ActiveUnit     uuid.UUID
	AwaitingUnits  map[uuid.UUID]*Unit
	ActivatedUnits map[uuid.UUID]*Unit
}

type Game struct {
	activePlayer  *Player
	passivePlayer *Player
}

func (g *Game) Run() chan Action {
	actionChan := make(chan Action)

	go func() {
		g.gameProcess(actionChan)
	}()

	return actionChan
}

func (g *Game) gameProcess(actionChan chan Action) {
	// Deployment Roll-off

	for i := 0; i < 4; i++ {
		// 1. Round Begin

		// 2. Turns
		for len(g.activePlayer.AwaitingUnits) != 0 || len(g.passivePlayer.AwaitingUnits) != 0 {
			if len(g.activePlayer.AwaitingUnits) != 0 {
				// Activate Unit
				g.activatePhase(actionChan)

				// Moving
				performedAction := g.movePhase(actionChan)

				if performedAction == "hold" || performedAction == "advance" {
					// Shooting
					g.shootingPhase(actionChan)
				} else if performedAction == "charge" {
					// Fighting
				}
			}
			g.switchPlayers()
		}

		// 3. Round End

	}
}

func (g *Game) switchPlayers() {
	tmp := g.activePlayer
	g.activePlayer = g.passivePlayer
	g.passivePlayer = tmp
}

func (g *Game) activatePhase(actionChan chan Action) {
	for {
		action := <-actionChan
		if action.PlayerID != g.activePlayer.ID {
			// write error response
			continue
		}

		if action.ActivateAction == nil {
			// write error response
			continue
		}

		_, ok := g.activePlayer.AwaitingUnits[action.ActivateAction.UnitID]
		if !ok {
			// write error response
			continue
		}

		g.activePlayer.ActiveUnit = action.ActivateAction.UnitID
		return
	}
}

func (g *Game) movePhase(actionChan chan Action) string {
	for {
		action := <-actionChan
		if action.PlayerID != g.activePlayer.ID {
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

func (g *Game) shootingPhase(actionChan chan Action) {
	for {
		action := <-actionChan
		if action.PlayerID != g.activePlayer.ID {
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

func (g *Game) fightingPhase(actionChan chan Action) {
	for {
		action := <-actionChan
		if action.PlayerID != g.activePlayer.ID {
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
