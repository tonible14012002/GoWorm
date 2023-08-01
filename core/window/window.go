package window

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/core/common"
)

type Window struct {
	Size  common.Vector
	Title string
	clock time.Time
}

func (window *Window) resetClock() float64 {
	elapsedTime := float64(time.Since(window.clock))
	window.clock = time.Now()
	return elapsedTime
}

func (window *Window) Setup() {
	// Constructor
	if window.Title == "" {
		window.Title = "GO Game"
	}
	if common.IsEqual(window.Size, common.Vector{X: 0, Y: 0}) {
		window.Size = common.Vector{X: 640, Y: 320}
	}
	ebiten.SetWindowTitle(window.Title)
}

func (window *Window) Update() error {
	elapsedTime := window.resetClock()
	// TODO: pass elapsed time to game object's update
	fmt.Println("elapsed time: ", elapsedTime)
	return nil
}

func (window *Window) Draw(screen *ebiten.Image) {
	// TODO: pass screen to game object's render method
	fmt.Println("Running Draw", screen)
}

func (window *Window) Layout(outerWidth, outerHeight int) (screenWidth, screenHeight int) {
	return window.Size.X, window.Size.Y
}
