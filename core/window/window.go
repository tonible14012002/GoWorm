package window

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/core/common"
	eventutil "github.com/tonible14012002/go_game/core/event"
)

type Window struct {
	Size     common.Vector
	Title    string
	clock    time.Time
	EManager eventutil.EventManager
}

func (window *Window) resetClock() float64 {
	elapsedTime := float64(time.Since(window.clock))
	window.clock = time.Now()
	return elapsedTime
}

func (window *Window) Setup() {
	window.EManager.Setup()
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
	window.EManager.Update(0)
	fmt.Println(elapsedTime)
	return nil
}

func (window *Window) Draw(screen *ebiten.Image) {
	// TODO: pass screen to game object's render metho
}

func (window *Window) Layout(outerWidth, outerHeight int) (screenWidth, screenHeight int) {
	return window.Size.X, window.Size.Y
}
