package state

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/event"
	"github.com/tonible14012002/go_game/engine/schema"
)

type BaseState interface {
	OnCreate(*StateManager, *event.EventManager)
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
	Statetype schema.StateType
	GameState BaseState
}

type StateStack []*StateInfo
type StateGenerator func() BaseState
type StateFactory map[schema.StateType]StateGenerator
