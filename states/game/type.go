package game

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
const PLAYER_ASSET_SRC = BASE_ASSET_SRC + "/sprites/player/"

var PlayerSpriteInfos = map[PlayerVariant]map[PlayerStatus]animation.SpriteInfo{
	BLUE: {
		IDLE: {
			Src:         PLAYER_ASSET_SRC + "B_witch_idle.png",
			ColumnCount: 1,
			RowCount:    6,
			TotalFrame:  6,
			FrameDir:    animation.DOWN,
		},
		DEATH: {
			Src:         PLAYER_ASSET_SRC + "B_witch_death.png",
			ColumnCount: 1,
			RowCount:    12,
			TotalFrame:  12,
			FrameDir:    animation.DOWN,
		},
	},
}

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
}

type Entities []EntityHandler
