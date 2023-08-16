package game

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/animation"
	"github.com/tonible14012002/go_game/engine/common"
)

const PLAYER_DEFAULT_SIZE = 10

type PlayerTeam struct {
	players        []*PlayerEntity
	activePlayerId int
}

func (t *PlayerTeam) CreatePlayer(size int, spriteInfo animation.SpriteInfo, info ...common.Vectorf) *PlayerEntity {
	newPlayer := createPlayer(size, spriteInfo, info...)
	t.players = append(t.players, newPlayer)
	return newPlayer
}

func (t *PlayerTeam) GetNextPlayer() *PlayerEntity {
	if len(t.players) == 0 {
		log.Fatal("must create team before get next player")
	}
	if t.IsAllDeatch() {
		return nil
	}
	activePlayerId := (t.activePlayerId + 1) % len(t.players)
	// for t.players[activePlayerId].IsDeath() {
	// 	activePlayerId = (activePlayerId + 1) % len(t.players)
	// 	fmt.Println("loop check", activePlayerId)
	// }
	t.activePlayerId = activePlayerId
	return t.players[activePlayerId]
}

func (t *PlayerTeam) GetAllPlayers() []*PlayerEntity {
	return t.players
}

func (t *PlayerTeam) UpdateTeam(elapsed time.Duration) {
	for i := range t.players {
		t.players[i].Update(elapsed)
	}
}

func (t *PlayerTeam) RenderTeam(screen *ebiten.Image) {
	for i := range t.players {
		t.players[i].Render(screen)
	}
}

func (t *PlayerTeam) IsAllStable() bool {
	for i := range t.players {
		if t.players[i].isStable {
			return false
		}
	}

	return true
}

func (t *PlayerTeam) IsAllDeatch() bool {
	for i := range t.players {
		if !t.players[i].IsDeath() {
			return false
		}
	}
	return true
}
