package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/core/state"
	"github.com/tonible14012002/go_game/core/window"
)

type Game struct {
	myWindow     window.Window
	stateManager state.StateManager
}

func (game *Game) Setup() {
	game.myWindow = window.Window{Title: "WORM GAME"}
	game.myWindow.Setup()
	game.stateManager.Setup()
}

func (game *Game) Start() error {
	return ebiten.RunGame(&game.myWindow)
}
