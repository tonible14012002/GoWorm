package game

import (
	"image/color"

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

func getMissileColor(value int) color.RGBA {
	if value < bufferMax/5 {
		return color.RGBA{0xef, 0x62, 0x62, 0xff}
	}
	if value < bufferMax*2/5 {
		return color.RGBA{0xf3, 0x99, 0x61, 0xff}
	}
	if value < bufferMax*3/5 {
		return color.RGBA{0xf7, 0xd0, 0x60, 0xff}
	}
	if value < bufferMax*4/5 {
		return color.RGBA{0x7c, 0xe8, 0x95, 0xff}
	}
	return color.RGBA{0x00, 0xff, 0xca, 0xff}
}
