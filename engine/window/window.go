package window

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/common"
	"github.com/tonible14012002/go_game/engine/event"
	"github.com/tonible14012002/go_game/engine/state"
	"github.com/tonible14012002/go_game/states/intro"
	"github.com/tonible14012002/go_game/states/menu"
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
	window.stateMgr.RegisterState(state.Intro, func() state.BaseState {
		return &intro.StateIntro{}
	})
	window.stateMgr.RegisterState(state.Menu, func() state.BaseState {
		return &menu.StateGame{}
	})
	window.EManager.AddCallback(state.Intro, "ENTER", func(*event.EventDetail) {
		fmt.Println("Switching to menu....")
		window.stateMgr.SwitchTo(state.Menu)
	})
	window.stateMgr.SwitchTo(state.Intro)

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
	window.stateMgr.Update(elapsedTime)
	window.EManager.Update(window.stateMgr.GetCurrentState())
	window.stateMgr.LateUpdate()
	return nil
}

func (window *Window) Draw(screen *ebiten.Image) {
	window.stateMgr.Render()
}

func (window *Window) Layout(outerWidth, outerHeight int) (screenWidth, screenHeight int) {
	return window.Size.X, window.Size.Y
}
