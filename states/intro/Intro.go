package intro

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/common"
	"github.com/tonible14012002/go_game/engine/event"
	"github.com/tonible14012002/go_game/engine/state"
	"github.com/tonible14012002/go_game/game/entity"
)

type StateIntro struct {
	isTransparent   bool
	isTranscendent  bool
	clock           time.Duration
	introTextSprite *entity.Player
}

func (intro *StateIntro) OnCreate(stateMgr *state.StateManager, eventMgr *event.EventManager) {
	intro.clock = 0
	x, _ := ebiten.WindowSize()
	intro.introTextSprite = &entity.Player{
		Pos:  common.Vectorf{X: float64(x) / 2, Y: 0},
		Velo: common.Vectorf{X: 0, Y: 20},
	}
	intro.introTextSprite.Setup()
	fmt.Println(intro.introTextSprite.Pos)
}

func (intro *StateIntro) OnDestroy() {
}

func (intro *StateIntro) Activate() {
}

func (intro *StateIntro) Deactivate() {
}

func (intro *StateIntro) Update(elapsed time.Duration) {
	fmt.Println("update state")
	intro.clock += elapsed
	if intro.clock.Seconds() >= 6 {
		intro.introTextSprite.Velo = common.Vectorf{} // Velo=0
	}
	intro.introTextSprite.Update(elapsed)
}

func (intro *StateIntro) Render(screen *ebiten.Image) {
	intro.introTextSprite.Render(screen)
}

func (intro *StateIntro) SetTransparent(isTransparent bool) {
	intro.isTransparent = isTransparent
}

func (intro *StateIntro) IsTransparent() bool {
	return intro.isTransparent
}

func (intro *StateIntro) SetTranscendent(isTranscendent bool) {
	intro.isTranscendent = isTranscendent
}

func (intro *StateIntro) IsTranscendent() bool {
	return intro.isTranscendent
}
