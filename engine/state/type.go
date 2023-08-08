package state

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type StateType int

const (
	Global StateType = -1
	Intro  StateType = iota
	Menu
	Game
	Ending
)

type BaseState interface {
	OnCreate()
	OnDestroy()
	Activate()
	Deactivate()
	Update(time.Duration)
	Render(*ebiten.Image)
	SetTransparent(bool)
	IsTransparent() bool
	SetTranscendent(bool)
	IsTranscendent() bool
}

type StateInfo struct {
	Statetype StateType
	GameState BaseState
}

type StateStack []*StateInfo
type StateGenerator func() BaseState
type StateFactory map[StateType]StateGenerator
