package game

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/common"
	"github.com/tonible14012002/go_game/engine/constant"
	"github.com/tonible14012002/go_game/engine/event"
	"github.com/tonible14012002/go_game/engine/schema"
	"github.com/tonible14012002/go_game/engine/state"
	"golang.org/x/exp/slices"
)

type StateGame struct {
	stateMgr *state.StateManager
	eventMgr *event.EventManager
	world    WorldMap
	entities Entities
	randGen  *rand.Rand
}

func (game *StateGame) OnCreate(stateMgr *state.StateManager, eventMgr *event.EventManager) {
	game.stateMgr = stateMgr
	game.eventMgr = eventMgr
	x, y := constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT
	game.world = createWorld(x/3, y/3, 3, common.Vectorf{
		X: 0,
		Y: 500,
	})

	seed := time.Now().Second()
	game.randGen = rand.New(rand.NewSource(int64(seed)))
}

func (game *StateGame) OnDestroy() {
}

func (game *StateGame) Activate() {
	game.eventMgr.AddCallback(schema.Game, "ESC", func(ed *event.EventDetail) {
		game.stateMgr.SwitchTo(schema.Intro)
	})
	game.eventMgr.AddCallback(schema.Game, "MouseLeftClick", game.AddEntityOnClick)
	game.eventMgr.AddCallback(schema.Game, "CtrlMouseLeftClick", func(ed *event.EventDetail) {
		fmt.Println("Boomb")
		game.Boom(common.Vectorf{X: float64(ed.MouseX), Y: float64(ed.MouseY)})
	})
}

func (game *StateGame) Deactivate() {
	game.eventMgr.RemoveCallback(schema.Game, "ESC")
	game.eventMgr.RemoveCallback(schema.Game, "MouseLeftClick")
	game.eventMgr.RemoveCallback(schema.Game, "CtrlMouseLeftClick")
}

func (game *StateGame) Update(elapsed time.Duration) {
	toRemoveEntityIndices := game.world.UpdatePhysic(elapsed, game.entities)
	remainEntities := make([]EntityHandler, 0, len(game.entities)-len(toRemoveEntityIndices))
	for i := range game.entities {
		if !slices.Contains(toRemoveEntityIndices, i) {
			remainEntities = append(remainEntities, game.entities[i])
		}
	}
	game.entities = remainEntities
}

func (game *StateGame) Render(screen *ebiten.Image) {
	game.world.Render(screen)
	for _, entity := range game.entities {
		entity.Render(screen)
	}
}

func (game *StateGame) IsTransparent() bool {
	return false
}

func (game *StateGame) IsTranscendent() bool {
	return false
}

func (game *StateGame) AddEntityOnClick(detail *event.EventDetail) {
	mousePos := common.Vectorf{
		X: float64(detail.MouseX),
		Y: float64(detail.MouseY),
	}
	newObject := EntityHandler(createPlayer(10, mousePos))
	game.entities = append(game.entities, newObject)
}

func (game *StateGame) Boom(mousePos common.Vectorf) {
	debrises := make([]EntityHandler, 20)
	for i := range debrises {
		debrisVelo := common.Vectorf{
			X: math.Cos(game.randGen.Float64()*2*math.Pi) * 100,
			Y: math.Sin(game.randGen.Float64()*2*math.Pi) * 100,
		}
		debrises[i] = EntityHandler(createObject(3, mousePos, debrisVelo))
	}
	game.entities = append(game.entities, debrises...)
}
