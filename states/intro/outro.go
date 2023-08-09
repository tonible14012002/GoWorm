package intro

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/tonible14012002/go_game/engine/event"
	"github.com/tonible14012002/go_game/engine/schema"
	"github.com/tonible14012002/go_game/engine/state"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type SectionType int

const (
	NAME SectionType = iota
	MESSAGE
)

type StateIntro struct {
	isTransparent  bool
	isTranscendent bool
	contents       map[SectionType]string
	fontFaces      map[SectionType]font.Face
	posXs          map[SectionType]float64
	posYs          map[SectionType]float64
	stateMgr       *state.StateManager
	eventMgr       *event.EventManager
}

func (intro *StateIntro) OnCreate(stateMgr *state.StateManager, eventMgr *event.EventManager) {
	intro.contents = make(map[SectionType]string)
	intro.fontFaces = make(map[SectionType]font.Face)
	intro.posXs = make(map[SectionType]float64)
	intro.posYs = make(map[SectionType]float64)

	intro.stateMgr = stateMgr
	intro.eventMgr = eventMgr

	x, y := ebiten.WindowSize()
	intro.contents[NAME] = "Go Worm"
	intro.contents[MESSAGE] = "Press 'A' to start"

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

	messageMplusNormalFont, messageFontFaceErr := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if messageFontFaceErr != nil {
		log.Fatal(messageFontFaceErr)
	}

	intro.fontFaces[NAME] = nameMplusNormalFont
	intro.fontFaces[MESSAGE] = messageMplusNormalFont

	nameBound := text.BoundString(intro.fontFaces[NAME], intro.contents[NAME])
	messageBound := text.BoundString(intro.fontFaces[MESSAGE], intro.contents[MESSAGE])

	intro.posXs[NAME] = (float64(x) - (float64(nameBound.Max.X) - float64(nameBound.Min.X))) / 2
	intro.posYs[NAME] = 0
	intro.posXs[MESSAGE] = (float64(x) - (float64(messageBound.Max.X) - float64(messageBound.Min.X))) / 2
	intro.posYs[MESSAGE] = float64(y)/2 - float64(nameBound.Min.Y)
}

func (intro *StateIntro) OnDestroy() {
}

func (intro *StateIntro) Activate() {
	intro.eventMgr.AddCallback(schema.Intro, "ENTER", func(ed *event.EventDetail) {
		intro.stateMgr.SwitchTo(schema.Game)
	})
}

func (intro *StateIntro) Deactivate() {
	intro.eventMgr.RemoveCallback(schema.Intro, "ENTER")
}

func (intro *StateIntro) Update(elapsed time.Duration) {
	fmt.Println("update state")

	_, y := ebiten.WindowSize()

	if intro.posYs[NAME] <= float64(y)/2 {
		intro.posYs[NAME] += float64(elapsed.Seconds() * 200)
	}
}

func (intro *StateIntro) Render(screen *ebiten.Image) {
	_, y := ebiten.WindowSize()

	text.Draw(screen, intro.contents[NAME], intro.fontFaces[NAME], int(intro.posXs[NAME]), int(intro.posYs[NAME]), color.RGBA{0x65, 0x28, 0xf7, 0xff})

	if intro.posYs[NAME] > float64(y)/2 {
		text.Draw(screen, intro.contents[MESSAGE], intro.fontFaces[MESSAGE], int(intro.posXs[MESSAGE]), int(intro.posYs[MESSAGE]), color.RGBA{0xff, 0xff, 0xff, 0xff})
	}
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
