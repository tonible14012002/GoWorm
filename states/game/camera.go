package game

import (
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/common"
)

type CameraDirection uint8

const (
	NONE CameraDirection = iota
	UP
	DOWN
	LEFT
	RIGHT
)

type Camera struct {
	ViewPort   common.Vectorf
	Pos        common.Vectorf
	zoomFactor float64
	rotation   float64
	Cam        *ebiten.Image
	direction  CameraDirection
	speed      float64
}

func (c *Camera) SetDirection(dir CameraDirection) {
	c.direction = dir
}

func (c *Camera) GetViewPortCenter() common.Vectorf {
	return common.Vectorf{
		X: c.ViewPort.X / 2,
		Y: c.ViewPort.Y / 2,
	}
}

func (c *Camera) getCameraGeoM() ebiten.GeoM {
	m := ebiten.GeoM{}
	viewPortCenter := c.GetViewPortCenter()
	m.Translate(-c.Pos.X-viewPortCenter.X, -c.Pos.Y-viewPortCenter.Y)
	// We want to scale and rotate around center of image / screen
	m.Scale(
		math.Pow(1.01, float64(c.zoomFactor)),
		math.Pow(1.01, float64(c.zoomFactor)),
	)
	m.Rotate(float64(c.rotation) * 2 * math.Pi / 360)
	m.Translate(viewPortCenter.X, viewPortCenter.Y)
	return m
}

func (c *Camera) GetRenderScreen() *ebiten.Image { return c.Cam }

func (c *Camera) Render(screen *ebiten.Image) {
	screen.DrawImage(c.Cam, &ebiten.DrawImageOptions{
		GeoM: c.getCameraGeoM(),
	})

	c.Cam.Fill(color.Black)
}

func (c *Camera) Update(elapsed time.Duration) {
	switch c.direction {
	case LEFT:
		{
			c.Pos = c.Pos.Add(common.Vectorf{X: -c.speed}.Multi(elapsed.Seconds()))
		}
	case RIGHT:
		{
			c.Pos = c.Pos.Add(common.Vectorf{X: c.speed}.Multi(elapsed.Seconds()))
		}
	case DOWN:
		{
			c.Pos = c.Pos.Add(common.Vectorf{Y: c.speed}.Multi(elapsed.Seconds()))
		}
	case UP:
		{
			c.Pos = c.Pos.Add(common.Vectorf{Y: -c.speed}.Multi(elapsed.Seconds()))
		}
	}
	c.LateUpdate()
}

func (c *Camera) SetCameraSpeed(speed float64) {
	c.speed = speed
}

func (c *Camera) LateUpdate() {
	c.direction = NONE
}

func (c *Camera) Move(dir CameraDirection) {
	c.direction = dir
}
