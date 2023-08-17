package game

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tonible14012002/go_game/engine/common"
)

type Object struct {
	pos              common.Vectorf
	velo             common.Vectorf
	accel            common.Vectorf
	isStable         bool
	radius           int
	boundBeforeDeath uint8
}

func (o *Object) Setup(radius int, info ...common.Vectorf) *Object {
	// info = pos, vel, accel
	o.radius = radius
	o.isStable = false
	switch len(info) {
	case 1:
		o.pos = info[0]
	case 2:
		o.pos = info[0]
		o.velo = info[1]
	case 3:
		o.pos = info[0]
		o.velo = info[1]
		o.accel = info[2]
	}
	return o
}

func (o *Object) GetRadius() int { return o.radius }

func (o *Object) GetPosition() common.Vectorf {
	return o.pos
}
func (o *Object) SetPosition(pos common.Vectorf) { o.pos = pos }

func (o *Object) GetVelo() common.Vectorf     { return o.velo }
func (o *Object) SetVelo(velo common.Vectorf) { o.velo = velo }

func (o *Object) GetAccel() common.Vectorf      { return o.accel }
func (o *Object) SetAccel(accel common.Vectorf) { o.accel = accel }

func (o *Object) IsStable() bool        { return o.isStable }
func (o *Object) SetStable(stable bool) { o.isStable = stable }

func (o *Object) Update(elapsed time.Duration) {

}

func (o *Object) GetFriction() float64 { return 0.5 }

func (o *Object) Render(screen *ebiten.Image) {
	posX := float32(o.pos.X)
	posY := float32(o.pos.Y)
	vector.DrawFilledCircle(screen, posX, posY, float32(o.radius), color.RGBA{0x27, 0x37, 0x4d, 0xff}, false)
}
func (o *Object) IsDeath() bool {
	return o.boundBeforeDeath > 2
}

func (o *Object) ToBeRemove() bool {
	return o.IsDeath()
}

func (o *Object) DoBouncing() {
	o.boundBeforeDeath++
}
func (o *Object) DoFalling() {}
func (o *Object) DoBomb()    {}
func (o *Object) IsExplosion() (bool, *common.Vectorf, int) {
	return false, nil, 0
}
