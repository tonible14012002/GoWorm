package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/event"
	"github.com/tonible14012002/go_game/engine/state"
)

type StateGame struct {
	isTransparent  bool
	isTranscendent bool
}

func (intro *StateGame) OnCreate(stateMgr *state.StateManager, eventMgr *event.EventManager) {
}

func (intro *StateGame) OnDestroy() {
}

func (intro *StateGame) Activate() {
}

func (intro *StateGame) Deactivate() {
}

func (intro *StateGame) Update(elapsed time.Duration) {
}

func (intro *StateGame) Render(screen *ebiten.Image) {
}

func (intro *StateGame) SetTransparent(isTransparent bool) {
}

func (intro *StateGame) IsTransparent() bool {
	return intro.isTransparent
}

func (intro *StateGame) SetTranscendent(isTranscendent bool) {
	intro.isTranscendent = isTranscendent
}

func (intro *StateGame) IsTranscendent() bool {
	return intro.isTranscendent
}
