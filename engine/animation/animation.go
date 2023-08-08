package animation

import (
	"fmt"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type FrameDirection uint8

const (
	RIGHT FrameDirection = iota
	DOWN
)

type Animation struct {
	Src            string
	size           image.Point
	TotalFrame     int
	currentFrame   int
	speed          float64 // frames per second
	clock          time.Duration
	RowCount       int
	ColumnCount    int
	FrameDir       FrameDirection
	subImgPos      image.Point
	subImgSize     image.Point
	img            *ebiten.Image
	PeriodDuration float64
}

func (a *Animation) GetSpriteSize() image.Point {
	return a.subImgSize
}

func (a *Animation) Setup() {
	a.clock = 0
	img, _, err := ebitenutil.NewImageFromFile(a.Src)
	if err != nil {
		log.Fatal(err, "asoidjaosidj")
	}
	a.img = img
	a.size = a.img.Bounds().Size()

	a.subImgSize = image.Point{
		X: a.size.X / a.ColumnCount,
		Y: a.size.Y / a.RowCount,
	}
	a.speed = float64(a.TotalFrame) / a.PeriodDuration
}

func (a *Animation) Update(elapsed time.Duration) {
	a.clock += elapsed
	if a.clock.Seconds() >= 1/a.speed {
		a.currentFrame++
		a.clock = 0
	}
	if a.currentFrame >= a.TotalFrame {
		a.currentFrame = 0
	}
	switch a.FrameDir {
	case DOWN:
		{
			a.subImgPos = image.Point{
				X: a.subImgSize.X * (a.currentFrame / a.RowCount),
				Y: a.subImgSize.Y * (a.currentFrame % a.RowCount),
			}
		}
	case RIGHT:
		{
			a.subImgPos = image.Point{
				X: a.subImgSize.X * (a.currentFrame % a.ColumnCount),
				Y: a.subImgSize.Y * (a.currentFrame / a.RowCount),
			}
		}
	}
	fmt.Println(a.subImgPos, a.subImgSize)
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
