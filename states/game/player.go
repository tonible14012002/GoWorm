package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/animation"
	"github.com/tonible14012002/go_game/engine/common"
)

type PlayerEntity struct {
	pos       common.Vectorf
	velo      common.Vectorf
	accel     common.Vectorf
	isStable  bool
	radius    int
	animation animation.Animation
	variant   PlayerVariant
	status    PlayerStatus
}

func (p *PlayerEntity) Setup(radius int, info ...common.Vectorf) *PlayerEntity {
	// info = pos, vel, accel
	p.radius = radius
	p.isStable = false
	switch len(info) {
	case 1:
		p.pos = info[0]
	case 2:
		p.pos = info[0]
		p.velo = info[1]
	case 3:
		p.pos = info[0]
		p.velo = info[1]
		p.accel = info[2]
	}

	p.animation = animation.Animation{
		Info:           PlayerSpriteInfos[p.variant][p.status],
		PeriodDuration: 1,
	}
	p.animation.Setup()
	p.animation.StartAnimation(animation.FOREVER)

	return p
}

func (p *PlayerEntity) SetStatus() {
}

func (p *PlayerEntity) GetRadius() int { return p.radius }

func (p *PlayerEntity) GetPosition() common.Vectorf {
	return p.pos
}
func (p *PlayerEntity) SetPosition(pos common.Vectorf) { p.pos = pos }

func (p *PlayerEntity) GetVelo() common.Vectorf     { return p.velo }
func (p *PlayerEntity) SetVelo(velo common.Vectorf) { p.velo = velo }

func (p *PlayerEntity) GetAccel() common.Vectorf      { return p.accel }
func (p *PlayerEntity) SetAccel(accel common.Vectorf) { p.accel = accel }

func (p *PlayerEntity) IsStable() bool        { return p.isStable }
func (p *PlayerEntity) SetStable(stable bool) { p.isStable = stable }

func (p *PlayerEntity) Update(elapsed time.Duration) {
	p.animation.Update(elapsed)
}

func (p *PlayerEntity) GetFriction() float64 { return 0.3 }

func (p *PlayerEntity) Render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	spriteSize := p.animation.GetSpriteSize()
	op.GeoM.Translate(p.pos.X-float64(spriteSize.X)/2, p.pos.Y-float64(spriteSize.Y)/2)
	p.animation.Render(screen, op)
}

func (p *PlayerEntity) IsDeath() bool {
	return false
}
func (p *PlayerEntity) DoBouncing() {}
func (p *PlayerEntity) DoFalling()  {}
func (p *PlayerEntity) DoBomb()     {}
func (p *PlayerEntity) ToBeRemove() bool {
	return false
}
