package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/event"
	"github.com/tonible14012002/go_game/engine/schema"
	"github.com/tonible14012002/go_game/engine/state"
)

type StateGame struct {
	stateMgr *state.StateManager
	eventMgr *event.EventManager
	world    WorldMap
}

func (game *StateGame) OnCreate(stateMgr *state.StateManager, eventMgr *event.EventManager) {
	game.stateMgr = stateMgr
	game.eventMgr = eventMgr
	x, y := ebiten.WindowSize()
	game.world = WorldMap{Width: x / 3, Height: y / 3, GraphicSize: 3}
	game.world.Setup()
}

func (game *StateGame) OnDestroy() {
}

func (game *StateGame) Activate() {
	game.eventMgr.AddCallback(schema.Game, "A", func(ed *event.EventDetail) {
		game.stateMgr.SwitchTo(schema.Intro)
	})
}

func (game *StateGame) Deactivate() {
	game.eventMgr.RemoveCallback(schema.Game, "A")

}

func (game *StateGame) Update(elapsed time.Duration) {
	game.world.Update(elapsed)
}

func (game *StateGame) Render(screen *ebiten.Image) {
	game.world.Render(screen)
}

func (game *StateGame) SetTransparent(isTransparent bool) {
}

func (game *StateGame) IsTransparent() bool {
	return false
}

func (game *StateGame) SetTranscendent(isTranscendent bool) {
}

func (game *StateGame) IsTranscendent() bool {
	return false
}
