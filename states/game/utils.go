package game

import (
	"image/color"
	"math"

	"github.com/tonible14012002/go_game/engine/animation"

	"github.com/tonible14012002/go_game/engine/common"
)

func createObject(size int, info ...common.Vectorf) *Object {
	return (&Object{}).Setup(size, info...)
}

func createWorld(width, height, graphicSize int, gravityAccel common.Vectorf) WorldMap {
	return *(&WorldMap{}).Setup(width, height, graphicSize, gravityAccel)
}

func createPlayer(size int, spriteInfo animation.SpriteInfo, info ...common.Vectorf) *PlayerEntity {
	return (&PlayerEntity{}).Setup(size, spriteInfo, info...)
}

func createMissile(posX, posY float64) *Missile {
	return (&Missile{}).Setup(common.Vectorf{X: posX, Y: posY})
}

func IsInsideCircle(x, y int, posX, posY float64, radius uint) bool {
	distanceSquared := math.Pow((float64(x)-posX), 2) + math.Pow((float64(y)-posY), 2)
	return distanceSquared <= math.Pow(float64(radius), 2)
}

func getMissileColor(value float64) color.RGBA {
	if value < maxEnergy/9 {
		return color.RGBA{0xef, 0x62, 0x62, 0xff}
	}
	if value < maxEnergy*2/9 {
		return color.RGBA{0xf0, 0x7c, 0x58, 0xff}
	}
	if value < maxEnergy*3/9 {
		return color.RGBA{0xf0, 0x96, 0x4f, 0xff}
	}
	if value < maxEnergy*4/9 {
		return color.RGBA{0xf1, 0xaf, 0x45, 0xff}
	}
	if value < maxEnergy*5/9 {
		return color.RGBA{0xf1, 0xc9, 0x3b, 0xff}
	}
	if value < maxEnergy*6/9 {
		return color.RGBA{0xb6, 0xd6, 0x55, 0xff}
	}
	if value < maxEnergy*7/9 {
		return color.RGBA{0x7a, 0xe2, 0x6f, 0xff}
	}
	if value < maxEnergy*8/9 {
		return color.RGBA{0x3f, 0xef, 0x88, 0xff}
	}
	return color.RGBA{0x03, 0xfb, 0xa2, 0xff}
}
