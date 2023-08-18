package game

import (
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tonible14012002/go_game/engine/common"
)

type Missile struct {
	pos               common.Vectorf
	velo              common.Vectorf
	accel             common.Vectorf
	bounceBeforeDeath int
	maxDamage         float64
}

func (m *Missile) Setup(pos common.Vectorf) *Missile {
	m.pos = pos
	m.velo.X = 0
	m.velo.Y = 0
	m.maxDamage = 60

	return m
}

func (m *Missile) GetPosition() common.Vectorf {
	return m.pos
}
func (m *Missile) SetPosition(pos common.Vectorf) { m.pos = pos }

func (m *Missile) GetVelo() common.Vectorf     { return m.velo }
func (m *Missile) SetVelo(velo common.Vectorf) { m.velo = velo }

func (m *Missile) GetAccel() common.Vectorf      { return m.accel }
func (m *Missile) SetAccel(accel common.Vectorf) { m.accel = accel }

func (m *Missile) Update(elapsed time.Duration) {}

func (m *Missile) Render(screen *ebiten.Image) {
	m.RenderMissile(screen)
}

func (m *Missile) RenderMissile(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(m.pos.X), float32(m.pos.Y), 5, 5, color.White, false)
}

func (m *Missile) Fire(angle float64, buffer float64) {
	mag := buffer
	m.velo = common.Vectorf{
		X: mag * math.Cos(angle) * 700,
		Y: mag * -math.Sin(angle) * 700,
	}
}
func (m *Missile) IsDeath() bool {
	return false
}
func (m *Missile) DoBouncing() {
	m.bounceBeforeDeath++
}
func (m *Missile) DoFalling() {}
func (m *Missile) DoBomb()    {}
func (m *Missile) ToBeRemove() bool {
	return m.bounceBeforeDeath >= 1
}
func (m *Missile) GetFriction() float64 { return 0.3 }

func (m *Missile) GetRadius() int { return 5 }
func (m *Missile) IsStable() bool { return true }
func (m *Missile) SetStable(bool) {}

func (m *Missile) IsExplosion() (bool, *common.Vectorf, int, float64) {
	return m.bounceBeforeDeath >= 1, &m.pos, 20, m.maxDamage
}
