package game

import (
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tonible14012002/go_game/engine/animation"
	"github.com/tonible14012002/go_game/engine/common"
)

type MovingDirection int

const (
	None MovingDirection = iota
	Up
	Down
)

const (
	crosshairAngleVelo float64 = 4
	crosshairRadius    float64 = 35
	crosshairScale     float64 = 0.3

	maxEnergy         float64 = 40
	maxChargingTime   float64 = 2
	maxHealth         float64 = 200
	healthRenderScale float64 = 4
	maxDamage         float64 = 60
)

type PlayerEntity struct {
	pos                common.Vectorf
	velo               common.Vectorf
	accel              common.Vectorf
	isStable           bool
	radius             int
	animation          animation.Animation
	crosshairSprite    animation.Animation
	crosshairAngle     float64
	isActive           bool
	health             float64
	energy             float64
	isCharing          bool
	crossHairDirection MovingDirection
}

func (p *PlayerEntity) Setup(radius int, spriteInfo animation.SpriteInfo, info ...common.Vectorf) *PlayerEntity {
	// info = pos, vel, accel
	p.radius = radius
	p.isStable = false
	p.crosshairAngle = 0
	p.crossHairDirection = None
	p.health = maxHealth

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
		Info: spriteInfo,
	}
	p.crosshairSprite = animation.Animation{
		Info: animation.SpriteInfo{
			Src:         "assets/sprites/crosshairs/crosshair.png",
			ColumnCount: 1,
			RowCount:    1,
			TotalFrame:  1,
		},
	}
	p.crosshairSprite.Setup()
	p.animation.Setup()
	p.animation.StartAnimation(animation.FOREVER)
	p.isCharing = false
	return p
}

func (p *PlayerEntity) SetStatus() {}

func (p *PlayerEntity) StartCharging() {
	p.isCharing = true
}

func (p *PlayerEntity) FireMissile() {
	p.isCharing = false
}

func (p *PlayerEntity) GetEnergyAmountPercent() float64 {
	return p.energy / maxChargingTime
}

func (p *PlayerEntity) GetEnergyAmountCharged() float64 {
	return p.energy * maxEnergy / maxChargingTime
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

	if p.isCharing {
		p.energy += float64(elapsed.Seconds()) * 2
	} else {
		p.energy = 0
	}
	if p.energy > maxChargingTime {
		p.energy = maxChargingTime
	}

	switch p.crossHairDirection {
	case Up:
		{
			p.crosshairAngle += elapsed.Seconds() * crosshairAngleVelo
		}
	case Down:
		{
			p.crosshairAngle -= elapsed.Seconds() * crosshairAngleVelo
		}
	}
	p.crossHairDirection = None
}

func (p *PlayerEntity) GetFriction() float64 { return 0.2 }

func (p *PlayerEntity) RenderCrosshair(screen *ebiten.Image) {
	if p.crosshairAngle > 2*math.Pi {
		p.crosshairAngle = 0
	}

	if p.crosshairAngle < 0 {
		p.crosshairAngle = 2 * math.Pi
	}

	offsetX := math.Cos(p.crosshairAngle) * crosshairRadius
	offsetY := -math.Sin(p.crosshairAngle) * crosshairRadius

	op := ebiten.DrawImageOptions{}
	size := p.crosshairSprite.GetSpriteSize()

	spriteX := p.pos.X + offsetX - (float64(size.X) / 2 * crosshairScale)
	spriteY := p.pos.Y + offsetY - (float64(size.Y) / 2 * crosshairScale)

	op.GeoM.Scale(crosshairScale, crosshairScale)
	op.GeoM.Translate(spriteX, spriteY)
	p.crosshairSprite.Render(screen, &op)
}

func (p *PlayerEntity) Render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	spriteSize := p.animation.GetSpriteSize()
	op.GeoM.Translate(p.pos.X-float64(spriteSize.X)/2, p.pos.Y-float64(spriteSize.Y)/2)
	p.animation.Render(screen, op)

	p.RenderHealth(screen)

	if p.isActive {
		p.RenderCrosshair(screen)
	}
	if p.isCharing {
		p.RenderMissileBuffer(screen)
	}
}

func (p *PlayerEntity) IsDeath() bool {
	return false
}
func (p *PlayerEntity) DoBouncing() {}
func (p *PlayerEntity) DoFalling()  {}
func (p *PlayerEntity) DoBomb(intArray []int) {
	originX := intArray[0]
	originY := intArray[1]
	radius := intArray[2]
	graphicSize := intArray[3]

	scaledOriginX := int(originX) / graphicSize
	scaledOriginY := int(originY) / graphicSize
	scaledPlayerX := int(p.pos.X) / graphicSize
	scaledPlayerY := int(p.pos.Y) / graphicSize

	distanceSquared := math.Pow((float64(scaledPlayerX)-float64(scaledOriginX)), 2) + math.Pow((float64(scaledPlayerY)-float64(scaledOriginY)), 2)

	if distanceSquared <= math.Pow(float64(radius), 2) {
		p.health = p.health - (1-distanceSquared/math.Pow(float64(radius), 2))*maxDamage
	}
}
func (p *PlayerEntity) ToBeRemove() bool {
	return false
}
func (p *PlayerEntity) SetIsActive(active bool) {
	p.isActive = active
}

func (p *PlayerEntity) SetMovingDirection(movingDirection MovingDirection) {
	p.crossHairDirection = movingDirection
}

func (p *PlayerEntity) RenderMissileBuffer(screen *ebiten.Image) {
	posX := float32(p.pos.X) - float32(p.animation.GetSpriteSize().X)/2 - 20
	posY := float32(p.pos.Y) + float32(maxEnergy)/2

	for i := 0; i <= int(maxEnergy)+1; i++ {
		if i < int(p.GetEnergyAmountCharged())+1 || i == 0 || i == int(maxEnergy)+1 {
			vector.DrawFilledRect(screen, posX, posY-float32(i), 7, 1, getMissileColor(p.GetEnergyAmountCharged()), false)

		} else {
			vector.DrawFilledRect(screen, posX, posY-float32(i), 1, 1, getMissileColor(p.GetEnergyAmountCharged()), false)
			vector.DrawFilledRect(screen, posX+6, posY-float32(i), 1, 1, getMissileColor(p.GetEnergyAmountCharged()), false)
		}
	}
}

func (p *PlayerEntity) IsExplosion() (bool, *common.Vectorf, int, float64) {
	return false, nil, 0, 0
}

func (p *PlayerEntity) RenderHealth(screen *ebiten.Image) {
	posX := float32(p.pos.X) - float32(maxHealth/healthRenderScale)/2
	posY := float32(p.pos.Y) - float32(p.animation.GetSpriteSize().Y)/2 - 20

	for i := 0; i <= int(maxHealth/healthRenderScale)+1; i++ {
		if i < int(p.health/healthRenderScale)+1 || i == 0 || i == int(maxHealth/healthRenderScale)+1 {
			vector.DrawFilledRect(screen, posX+float32(i), posY, 1, 5, color.RGBA{0xff, 0x00, 0x00, 0xff}, false)

		} else {
			vector.DrawFilledRect(screen, posX+float32(i), posY, 1, 1, color.RGBA{0xff, 0x00, 0x00, 0xff}, false)
			vector.DrawFilledRect(screen, posX+float32(i), posY+4, 1, 1, color.RGBA{0xff, 0x00, 0x00, 0xff}, false)
		}
	}
}
