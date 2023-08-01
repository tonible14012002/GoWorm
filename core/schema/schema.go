package schema

import "github.com/hajimehoshi/ebiten/v2"

type GameObject interface {
	Update(elapsedTime float64)
	Render(screen *ebiten.Image)
}
