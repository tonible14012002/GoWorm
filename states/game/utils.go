package game

import "github.com/tonible14012002/go_game/engine/common"

func createObject(size int, info ...common.Vectorf) *Object {
	return (&Object{}).Setup(size, info...)
}

func createWorld(width, height, graphicSize int, gravityAccel common.Vectorf) WorldMap {
	return *(&WorldMap{}).Setup(width, height, graphicSize, gravityAccel)
}

func createPlayer(size int, info ...common.Vectorf) *PlayerEntity {
	return (&PlayerEntity{}).Setup(size, info...)
}
