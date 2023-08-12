package gamemanager

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/common"
	"github.com/tonible14012002/go_game/engine/constant"
	"github.com/tonible14012002/go_game/engine/window"
)

type Game struct {
	myWindow window.Window
}

func (game *Game) Setup() {
	game.myWindow = window.Window{}
	game.myWindow.Setup("GOWORM", common.Vector{X: constant.SCREEN_WIDTH, Y: constant.SCREEN_HEIGHT})
}

func (game *Game) Start() error {
	return ebiten.RunGame(&game.myWindow)
}
