package window

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/common"
	"github.com/tonible14012002/go_game/engine/event"
	"github.com/tonible14012002/go_game/engine/schema"
	"github.com/tonible14012002/go_game/engine/state"
	"github.com/tonible14012002/go_game/states/game"
	"github.com/tonible14012002/go_game/states/intro"
)

type Window struct {
	Size     common.Vector
	Title    string
	clock    time.Time
	EManager event.EventManager
	stateMgr state.StateManager
}

func (window *Window) resetClock() time.Duration {
	elapsedTime := time.Since(window.clock)
	window.clock = time.Now()
	return elapsedTime
}

func (window *Window) Setup() {
	window.EManager.Setup()
	window.stateMgr.RegisterEventManager(&window.EManager)
	window.stateMgr.Setup()
	window.clock = time.Now()

	if window.Title == "" {
		window.Title = "GO Game"
	}
	if window.Size.IsEqual(common.Vector{X: 0, Y: 0}) {
		window.Size = common.Vector{X: 800, Y: 460}
	}
	ebiten.SetWindowSize(800, 460)
	ebiten.SetWindowTitle(window.Title)

	// Intro State
	window.stateMgr.RegisterState(schema.Intro, func() state.BaseState {
		return &intro.StateIntro{}
	})

	// Game State
	window.stateMgr.RegisterState(schema.Game, func() state.BaseState {
		return &game.StateGame{}
	})
	window.stateMgr.SwitchTo(schema.Intro)
}

func (window *Window) Update() error {
	elapsedTime := window.resetClock()
	window.stateMgr.Update(elapsedTime)
	window.EManager.Update(window.stateMgr.GetCurrentState())
	window.stateMgr.LateUpdate()
	return nil
}

func (window *Window) Draw(screen *ebiten.Image) {
	window.stateMgr.Render(screen)
}

func (window *Window) Layout(outerWidth, outerHeight int) (screenWidth, screenHeight int) {
	return window.Size.X, window.Size.Y
}
