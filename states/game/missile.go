package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tonible14012002/go_game/engine/common"
)

type Missile struct {
	pos      common.Vectorf
	velo     common.Vectorf
	accel    common.Vectorf
	buffer   int
	hasFired bool
}

var (
	bufferMax int = 100
)

func (m *Missile) Setup(pos common.Vectorf) *Missile {
	m.pos = pos
	m.velo.X = 0
	m.velo.Y = 0
	m.buffer = 0
	m.hasFired = false
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

func (m *Missile) Update(elapsed time.Duration) {
	if m.buffer < bufferMax && !m.hasFired {
		m.buffer = m.buffer + 1
	}
}

func (m *Missile) Render(screen *ebiten.Image) {
	posX := float32(m.pos.X)
	posY := float32(m.pos.Y)
	for i := 0; i < bufferMax; i++ {
		if i < m.buffer {
			vector.DrawFilledRect(screen, posX+float32(i), posY, 1, 5, getMissileColor(m.buffer), false)
		} else {
			vector.DrawFilledRect(screen, posX+float32(i), posY, 1, 1, getMissileColor(m.buffer), false)
			vector.DrawFilledRect(screen, posX+float32(i), posY+4, 1, 1, getMissileColor(m.buffer), false)
		}
	}
}

func (m *Missile) Fire() {
	m.hasFired = true
}
