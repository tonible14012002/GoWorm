package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/common"
)

type EntityHandler interface {
	GetPosition() common.Vectorf
	SetPosition(common.Vectorf)
	GetVelo() common.Vectorf
	SetVelo(common.Vectorf)
	GetAccel() common.Vectorf
	SetAccel(common.Vectorf)
	IsStable() bool
	SetStable(bool)
	Update(elapsed time.Duration)
	Render(screen *ebiten.Image)
	GetRadius() int
	GetFriction() float64
}

type Entities []EntityHandler
