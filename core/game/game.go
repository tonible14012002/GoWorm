package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
}

func (game *Game) Update() error {
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {

}

func (game *Game) Layout(outerWidth, outerHeight int) (screenWidth, screenHeight int) {
	return 640, 320
}
