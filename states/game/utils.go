package game

import "github.com/tonible14012002/go_game/engine/common"

func createObject(size int, pos common.Vectorf) *Object {
	return (&Object{}).Setup(size, pos)
}

func createWorld(width, height, graphicSize int, gravityAccel common.Vectorf) WorldMap {
	return *(&WorldMap{}).Setup(width, height, graphicSize, gravityAccel)
}
