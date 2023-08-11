package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/common"
	"github.com/tonible14012002/go_game/engine/event"
	"github.com/tonible14012002/go_game/engine/schema"
	"github.com/tonible14012002/go_game/engine/state"
)

type StateGame struct {
	stateMgr *state.StateManager
	eventMgr *event.EventManager
	world    WorldMap
	entities Entities
}

func (game *StateGame) OnCreate(stateMgr *state.StateManager, eventMgr *event.EventManager) {
	game.stateMgr = stateMgr
	game.eventMgr = eventMgr
	x, y := ebiten.WindowSize()
	game.world = createWorld(x/3, y/3, 3, common.Vectorf{
		X: 0,
		Y: 20,
	})
}

func (game *StateGame) OnDestroy() {
}

func (game *StateGame) Activate() {
	game.eventMgr.AddCallback(schema.Game, "ESC", func(ed *event.EventDetail) {
		game.stateMgr.SwitchTo(schema.Intro)
	})
	game.eventMgr.AddCallback(schema.Game, "MouseLeftClick", game.AddEntityOnClick)
}

func (game *StateGame) Deactivate() {
	game.eventMgr.RemoveCallback(schema.Game, "ESC")
	game.eventMgr.RemoveCallback(schema.Game, "MouseLeftClick")
}

func (game *StateGame) Update(elapsed time.Duration) {
	game.world.UpdatePhysic(elapsed, game.entities)
	game.world.Update(elapsed)
}

func (game *StateGame) Render(screen *ebiten.Image) {
	game.world.Render(screen)
	for _, entity := range game.entities {
		entity.Render(screen)
	}
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

func (game *StateGame) AddEntityOnClick(detail *event.EventDetail) {
	mousePos := common.Vectorf{
		X: float64(detail.MouseX),
		Y: float64(detail.MouseY),
	}
	newObject := EntityHandler(createObject(20, mousePos))
	game.entities = append(game.entities, newObject)
}
