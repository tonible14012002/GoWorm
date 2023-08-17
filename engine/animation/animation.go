package animation

import (
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tonible14012002/go_game/engine/common"
)

type FrameDirection uint8
type AnimationType uint8

const (
	FOREVER = iota
	SPECIFY
)

const (
	RIGHT FrameDirection = iota
	DOWN
)

type SpriteInfo struct {
	Src            string
	TotalFrame     int
	ColumnCount    int
	RowCount       int
	FrameDir       FrameDirection
	PeriodDuration float64
}

type Animation struct {
	Info         SpriteInfo
	size         common.Vector
	currentFrame int
	speed        float64 // frames per second
	clock        time.Duration
	subImgPos    common.Vector
	subImgSize   common.Vector
	img          *ebiten.Image
	enable       bool
	aType        AnimationType
	loopCount    int
	maxLoopCount int
}

func (a *Animation) GetSpriteSize() common.Vector {
	return a.subImgSize
}

func (a *Animation) StartAnimation(atype AnimationType, loopCount ...int) {
	a.clock = 0
	a.enable = true
	a.aType = atype
	a.loopCount = 0
	if len(loopCount) > 1 {
		log.Fatal("Just provide an integer")
	}
	if atype == SPECIFY {
		if len(loopCount) == 0 {
			log.Fatal("Must specify loopCount")
		}
		a.maxLoopCount = loopCount[0]
	}
}

func (a *Animation) StopAnimation() {
	a.clock = 0
	a.enable = false
	a.loopCount = 0
}

func (a *Animation) Setup() {
	a.clock = 0
	img, _, err := ebitenutil.NewImageFromFile(a.Info.Src)
	if err != nil {
		log.Fatal(err, "Setup Animation fail")
	}
	a.img = img
	size := a.img.Bounds().Size()
	a.size = common.Vector{
		X: size.X,
		Y: size.Y,
	}

	a.subImgSize = common.Vector{
		X: a.size.X / a.Info.ColumnCount,
		Y: a.size.Y / a.Info.RowCount,
	}
	a.speed = float64(a.Info.TotalFrame) / a.Info.PeriodDuration
}

func (a *Animation) UpdateInfo(info SpriteInfo) {
	a.Info = info
}

func (a *Animation) Reset() {
	a.Setup()
}

func (a *Animation) Update(elapsed time.Duration) {
	if !a.enable {
		return
	}

	if a.aType == SPECIFY && a.loopCount == a.maxLoopCount {
		return
	}
	if a.aType == SPECIFY {
		a.loopCount++
	}

	a.clock += elapsed
	if a.clock.Seconds() >= 1/a.speed {
		a.currentFrame++
		a.clock = 0
	}
	if a.currentFrame >= a.Info.TotalFrame {
		a.currentFrame = 0
	}
	switch a.Info.FrameDir {
	case DOWN:
		{
			a.subImgPos = common.Vector{
				X: a.subImgSize.X * (a.currentFrame / a.Info.RowCount),
				Y: a.subImgSize.Y * (a.currentFrame % a.Info.RowCount),
			}
		}
	case RIGHT:
		{
			a.subImgPos = common.Vector{
				X: a.subImgSize.X * (a.currentFrame % a.Info.ColumnCount),
				Y: a.subImgSize.Y * (a.currentFrame / a.Info.RowCount),
			}
		}
	}
}

func (a *Animation) Render(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	imgToRender := a.img.SubImage(image.Rect(
		a.subImgPos.X,
		a.subImgPos.Y,
		a.subImgPos.X+a.subImgSize.X,
		a.subImgPos.Y+a.subImgSize.Y,
	)).(*ebiten.Image)
	screen.DrawImage(imgToRender, op)
}
