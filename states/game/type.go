package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/common"
)

type PlayerStatus int
type PlayerVariant int
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
	IsDeath() bool
	DoBouncing()
	DoFalling()
	DoBomb()
	ToBeRemove() bool
	IsExplosion() (bool, *common.Vectorf, int, float64)
}

type Entities []EntityHandler
