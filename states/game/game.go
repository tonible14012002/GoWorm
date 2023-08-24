package game

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/tonible14012002/go_game/engine/animation"
	"github.com/tonible14012002/go_game/engine/common"
	"github.com/tonible14012002/go_game/engine/constant"
	"github.com/tonible14012002/go_game/engine/event"
	"github.com/tonible14012002/go_game/engine/schema"
	"github.com/tonible14012002/go_game/engine/state"
	"golang.org/x/exp/slices"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	teamCount          int = 2
	teamMemCount       int = 2
	maxStableThreshold int = 300
)

type StateGame struct {
	stateMgr        *state.StateManager
	eventMgr        *event.EventManager
	world           WorldMap
	entities        Entities
	randGen         *rand.Rand
	camera          Camera
	currentPlayer   *PlayerEntity
	playerTeams     []PlayerTeam
	currentTeamId   int
	teamCount       int
	teamMemCount    int
	gamePlayState   GamePlayState
	isStable        bool
	stableThreshold int
	isOver          bool
}

func (game *StateGame) OnCreate(stateMgr *state.StateManager, eventMgr *event.EventManager) {
	game.stateMgr = stateMgr
	game.eventMgr = eventMgr
	x, y := constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT
	game.world = createWorld(x/2, y/2, 2, common.Vectorf{
		X: 0,
		Y: 5,
	})

	seed := time.Now().Second()
	game.randGen = rand.New(rand.NewSource(int64(seed)))
	game.camera = Camera{
		ViewPort:   common.Vectorf{X: 800, Y: 400},
		Pos:        common.Vectorf{X: 0, Y: 0},
		zoomFactor: 50,
		rotation:   0,
		Cam:        ebiten.NewImage(constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT),
	}
	game.camera.SetCameraSpeed(300)

	game.teamCount = teamCount
	game.teamMemCount = teamMemCount
	game.playerTeams = make([]PlayerTeam, game.teamCount)
	for i := range game.playerTeams {
		team := &game.playerTeams[i]
		for j := 0; j < game.teamMemCount; j++ {
			var player *PlayerEntity
			if (i % game.teamCount) == 0 {
				player = team.CreatePlayer(
					PLAYER_DEFAULT_SIZE,
					animation.SpriteInfo{
						Src:            "assets/sprites/player/ChikBoy_idle.png",
						RowCount:       6,
						ColumnCount:    1,
						TotalFrame:     6,
						FrameDir:       animation.DOWN,
						PeriodDuration: 1,
					},
					common.Vectorf{
						X: (game.randGen.Float64())*float64(constant.SCREEN_WIDTH)*0.8 + 50,
						Y: 0,
					},
				)
			} else {
				player = team.CreatePlayer(
					PLAYER_DEFAULT_SIZE,
					animation.SpriteInfo{
						Src:            "assets/sprites/player/B_witch_idle.png",
						RowCount:       6,
						ColumnCount:    1,
						TotalFrame:     6,
						FrameDir:       animation.DOWN,
						PeriodDuration: 1,
					},
					common.Vectorf{
						X: (game.randGen.Float64()*0.8 + 0.2) * float64(constant.SCREEN_WIDTH),
						Y: 0,
					},
				)
			}
			game.entities = append(game.entities, player)
		}
	}
	game.currentTeamId = 0
	game.currentPlayer = game.playerTeams[game.currentTeamId].GetNextPlayer()
	game.currentPlayer.SetIsActive(true)
	game.gamePlayState.state = STANDBY
}

func (game *StateGame) OnDestroy() {
}

func (game *StateGame) Activate() {
	game.eventMgr.AddCallback(schema.Game, "ESC", func(ed *event.EventDetail) {
		game.stateMgr.SwitchTo(schema.Intro)
	})
	game.eventMgr.AddCallback(schema.Game, "SPACE", func(ed *event.EventDetail) {
		game.stateMgr.SwitchTo(schema.Menu)
	})
	game.eventMgr.AddCallback(schema.Game, "CtrlMouseLeftClick", func(ed *event.EventDetail) {
		game.Boom(common.Vectorf{X: float64(ed.MouseX), Y: float64(ed.MouseY)})
	})
	game.eventMgr.AddCallback(schema.Game, "ShiftArrowUp", func(ed *event.EventDetail) {
		game.camera.Move(UP)
	})
	game.eventMgr.AddCallback(schema.Game, "ShiftArrowDown", func(ed *event.EventDetail) {
		game.camera.Move(DOWN)
	})
	game.eventMgr.AddCallback(schema.Game, "ShiftArrowLeft", func(ed *event.EventDetail) {
		game.camera.Move(LEFT)
	})
	game.eventMgr.AddCallback(schema.Game, "ShiftArrowRight", func(ed *event.EventDetail) {
		game.camera.Move(RIGHT)
	})
	// game.eventMgr.AddCallback(schema.Game, "KeyN", func(ed *event.EventDetail) { game.NextPlayer() })
	game.eventMgr.AddCallback(schema.Game, "Comma", func(ed *event.EventDetail) { game.MoveCrosshair(Up) })
	game.eventMgr.AddCallback(schema.Game, "Dot", func(ed *event.EventDetail) { game.MoveCrosshair(Down) })
	game.eventMgr.AddCallback(schema.Game, "KeyXDown", func(ed *event.EventDetail) {
		if game.currentPlayer.isActive {
			game.InitMissile()
		}
	})
	game.eventMgr.AddCallback(schema.Game, "KeyXUp", func(ed *event.EventDetail) {
		if game.currentPlayer.isActive {
			game.FireMissile()
		}
	})
}

func (game *StateGame) Deactivate() {
	game.eventMgr.RemoveCallback(schema.Game, "ESC")
	game.eventMgr.RemoveCallback(schema.Game, "SPACE")
	game.eventMgr.RemoveCallback(schema.Game, "CtrlMouseLeftClick")
	game.eventMgr.RemoveCallback(schema.Game, "ShiftArrowUp")
	game.eventMgr.RemoveCallback(schema.Game, "ShiftArrowLeft")
	game.eventMgr.RemoveCallback(schema.Game, "ShiftArrowDown")
	game.eventMgr.RemoveCallback(schema.Game, "ShiftArrowRight")
	// game.eventMgr.RemoveCallback(schema.Game, "KeyN")
	game.eventMgr.RemoveCallback(schema.Game, "Comma")
	game.eventMgr.RemoveCallback(schema.Game, "Dot")
	game.eventMgr.RemoveCallback(schema.Game, "KeyXDown")
	game.eventMgr.RemoveCallback(schema.Game, "KeyXUp")
}

func (game *StateGame) Update(elapsed time.Duration) {
	// recalculate 10 times for correctiveness.
	for i := 0; i < 10; i++ {
		toRemoveEntityIndices, boomPoss := game.world.UpdatePhysic(elapsed, game.entities)

		remainEntities := make([]EntityHandler, 0, len(game.entities)-len(toRemoveEntityIndices))
		for i := range game.entities {
			if !slices.Contains(toRemoveEntityIndices, i) {
				remainEntities = append(remainEntities, game.entities[i])
			}
		}

		game.entities = remainEntities

		for i := range boomPoss {
			game.Boom(boomPoss[i])
		}
	}

	// UpdateTeam
	for i := range game.playerTeams {
		game.playerTeams[i].UpdateTeam(elapsed)
	}

	// Update Entity
	for i := range game.entities {
		game.entities[i].Update(elapsed)
	}

	if game.gamePlayState.state == FIRING || game.gamePlayState.state == EXPLODING {
		game.stableThreshold += 1
	}

	// Update game play state
	if (!game.IsGameStable() || game.gamePlayState.state == FIRING) && game.stableThreshold <= maxStableThreshold {
		game.isStable = false
	} else {
		game.isStable = true
	}
	if game.IsGameOver() {
		game.isOver = true
	} else {
		if game.gamePlayState.state == EXPLODING && game.isStable {
			game.gamePlayState.state = STANDBY
			game.stableThreshold = 0
			game.NextPlayer()
		}
	}

	// game.camera.Update(elapsed)
}

func (game *StateGame) Render(screen *ebiten.Image) {
	// game.world.Render(game.camera.Cam)
	game.world.Render(screen)
	for _, entity := range game.entities {
		// entity.Render(game.camera.Cam)
		entity.Render(screen)
	}

	if game.isOver {
		game.RenderGameOver(screen)
	}

	// game.camera.Render(screen)
	// }
}

func (game *StateGame) IsTransparent() bool {
	return false
}

func (game *StateGame) IsTranscendent() bool {
	return false
}

func (game *StateGame) Boom(pos common.Vectorf) {
	debrises := make([]EntityHandler, 20)
	for i := range debrises {
		debrisVelo := common.Vectorf{
			X: math.Cos(game.randGen.Float64()*2*math.Pi) * 20,
			Y: math.Sin(game.randGen.Float64()*2*math.Pi) * 20,
		}
		debrises[i] = EntityHandler(createObject(2, pos, debrisVelo))
	}
	game.entities = append(game.entities, debrises...)
	game.gamePlayState.state = EXPLODING
}

func (game *StateGame) IsGameStable() bool {
	for i := range game.entities {
		if !game.entities[i].IsStable() {
			return false
		}
	}
	return true
}

func (game *StateGame) IsGameOver() bool {
	for i := range game.playerTeams {
		team := &game.playerTeams[i]
		aliveCount := 0
		for j := 0; j < game.teamMemCount; j++ {
			if !team.players[j].IsDeath() {
				aliveCount++
			}
		}

		if aliveCount == 0 {
			return true
		}
	}

	return false
}

func (game *StateGame) MoveCrosshair(direction MovingDirection) {
	if game.currentPlayer != nil {
		game.currentPlayer.SetMovingDirection(direction)
	}
}

func (game *StateGame) NextPlayer() {
	foundNext := false
	for !foundNext {
		game.currentTeamId = (game.currentTeamId + 1) % game.teamCount
		game.currentPlayer = game.playerTeams[game.currentTeamId].GetNextPlayer()

		// skip death players
		if !game.currentPlayer.IsDeath() {
			foundNext = true
		}
	}
	game.currentPlayer.SetIsActive(true)
}

func (game *StateGame) InitMissile() {
	game.currentPlayer.StartCharging()
}

func (game *StateGame) FireMissile() {
	game.currentPlayer.FireMissile()
	missle := createMissile(game.currentPlayer.pos.X, game.currentPlayer.pos.Y)

	missle.Fire(game.currentPlayer.crosshairAngle, game.currentPlayer.GetEnergyAmountPercent())
	game.entities = append(game.entities, EntityHandler(missle))
	game.currentPlayer.SetIsActive(false)
	game.gamePlayState.state = FIRING
}

func (game *StateGame) RenderGameOver(screen *ebiten.Image) {
	x, y := ebiten.WindowSize()
	content := "Game over"

	tt, ttErr := opentype.Parse(fonts.PressStart2P_ttf)
	if ttErr != nil {
		log.Fatal(ttErr)
	}

	nameMplusNormalFont, nameFontFaceErr := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if nameFontFaceErr != nil {
		log.Fatal(nameFontFaceErr)
	}

	nameBound := text.BoundString(nameMplusNormalFont, content)

	posX := (float64(x) - (float64(nameBound.Max.X) - float64(nameBound.Min.X))) / 2
	posY := float64(y) / 2

	text.Draw(screen, content, nameMplusNormalFont, int(posX)-3, int(posY), color.RGBA{0xc4, 0xdd, 0xff, 0xff})
	text.Draw(screen, content, nameMplusNormalFont, int(posX), int(posY), color.RGBA{0x65, 0x28, 0xf7, 0xff})
}
