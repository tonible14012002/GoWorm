package entity

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/animation"
	"github.com/tonible14012002/go_game/engine/common"
)

type PlayerStatus int
type PlayerVariant int

const (
	IDLE PlayerStatus = iota
	ATTACK
	CHARGE
	DEATH
	RUN
	TAKE_DAMAGE
)

const (
	BLUE PlayerVariant = iota
)

const BASE_ASSET_SRC string = "assets"

var PlayerSpriteVariant = map[PlayerVariant]map[PlayerStatus]string{
	BLUE: {
		IDLE:        "B_witch_idle.png",
		ATTACK:      "B_witch_attack.png",
		CHARGE:      "B_witch_charge.png",
		DEATH:       "B_witch_death.png",
		RUN:         "B_witch_run.png",
		TAKE_DAMAGE: "B_witch_take_damage.png",
	},
}

func GetPlayerSpriteSrc(variant PlayerVariant, status PlayerStatus) string {
	return BASE_ASSET_SRC + "/sprites/player/" + PlayerSpriteVariant[variant][status]
}

type Player struct {
	Radius    int
	Pos       common.Vectorf
	Velo      common.Vectorf
	Accel     common.Vectorf
	Status    PlayerStatus
	Variant   PlayerVariant
	animation animation.Animation
}

func (p *Player) Setup() {
	p.Status = RUN
	p.Variant = BLUE
	p.animation = animation.Animation{
		Src:            GetPlayerSpriteSrc(p.Variant, p.Status),
		ColumnCount:    1,
		RowCount:       8,
		FrameDir:       animation.DOWN,
		PeriodDuration: 1,
		TotalFrame:     8,
	}
	p.animation.Setup()
}

func (p *Player) Update(elapsed time.Duration) {
	p.Pos.X += elapsed.Seconds() * p.Velo.X
	p.Pos.Y += elapsed.Seconds() * p.Velo.Y
	p.animation.Update(elapsed)
}

func (p *Player) Render(screen *ebiten.Image) {
	renderOp := &ebiten.DrawImageOptions{}
	renderOp.GeoM.Translate(p.Pos.X, p.Pos.Y)
	op := &ebiten.DrawImageOptions{}
	spriteSize := p.animation.GetSpriteSize()
	op.GeoM.Translate(p.Pos.X-float64(spriteSize.X)/2, p.Pos.Y-float64(spriteSize.Y)/2)
	p.animation.Render(screen, op)
}
