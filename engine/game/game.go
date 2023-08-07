package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/window"
)

type Game struct {
	myWindow window.Window
}

func (game *Game) Setup() {
	game.myWindow = window.Window{Title: "WORM GAME"}
	game.myWindow.Setup()
}

func (game *Game) Start() error {
	return ebiten.RunGame(&game.myWindow)
}
