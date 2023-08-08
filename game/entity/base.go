package entity

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameEntity interface {
	Setup()
	Update(elapsed time.Duration)
	Render(*ebiten.Image)
}
