package window

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/core/common"
	"github.com/tonible14012002/go_game/core/event"
	"github.com/tonible14012002/go_game/core/state"
)

type Window struct {
	Size     common.Vector
	Title    string
	clock    time.Time
	EManager event.EventManager
	stateMgr state.StateManager
}

func (window *Window) resetClock() float64 {
	elapsedTime := float64(time.Since(window.clock))
	window.clock = time.Now()
	return elapsedTime
}

func (window *Window) Setup() {
	window.EManager.Setup()
	window.stateMgr.Setup()
	if window.Title == "" {
		window.Title = "GO Game"
	}
	if window.Size.IsEqual(common.Vector{X: 0, Y: 0}) {
		window.Size = common.Vector{X: 640, Y: 320}
	}
	ebiten.SetWindowTitle(window.Title)
}

func (window *Window) Update() error {
	elapsedTime := window.resetClock()
	window.EManager.Update(window.stateMgr.GetCurrentState())
	window.stateMgr.Update(elapsedTime)
	window.stateMgr.LateUpdate()
	return nil
}

func (window *Window) Draw(screen *ebiten.Image) {
	window.stateMgr.Render()
}

func (window *Window) Layout(outerWidth, outerHeight int) (screenWidth, screenHeight int) {
	return window.Size.X, window.Size.Y
}
