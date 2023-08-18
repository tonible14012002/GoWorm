package menu

import (
	"image/color"
	"log"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tonible14012002/go_game/engine/event"
	"github.com/tonible14012002/go_game/engine/schema"
	"github.com/tonible14012002/go_game/engine/state"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/gomedium"
	"golang.org/x/image/font/opentype"
)

type SectionType int

const (
	NAME SectionType = iota
	MENU
	MENU_INNER
)

type MenuItemType int

const (
	RESUME MenuItemType = iota
	QUIT
	CONTROL
)

var (
	menuWidth  int = 550
	menuHeight int = 320
)

type StateMenu struct {
	fontFaces      map[SectionType]font.Face
	isMenuItemDown map[MenuItemType]bool
	stateMgr       *state.StateManager
	eventMgr       *event.EventManager
	showControl    bool
}

func (menu *StateMenu) OnCreate(stateMgr *state.StateManager, eventMgr *event.EventManager) {
	menu.fontFaces = make(map[SectionType]font.Face)
	menu.isMenuItemDown = make(map[MenuItemType]bool)

	menu.stateMgr = stateMgr
	menu.eventMgr = eventMgr

	ttPressStart, errPressStart := opentype.Parse(fonts.PressStart2P_ttf)
	if errPressStart != nil {
		log.Fatal(errPressStart)
	}

	ttBold, errBold := truetype.Parse(gobold.TTF)
	if errBold != nil {
		log.Fatal(errBold)
	}

	ttMedium, errMedium := truetype.Parse(gomedium.TTF)
	if errMedium != nil {
		log.Fatal(errMedium)
	}

	nameMplusNormalFont, nameFontFaceErr := opentype.NewFace(ttPressStart, &opentype.FaceOptions{
		Size:    48,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if nameFontFaceErr != nil {
		log.Fatal(nameFontFaceErr)
	}

	messageMplusNormalFont := truetype.NewFace(ttBold, &truetype.Options{
		Size:    20,
		DPI:     72,
		Hinting: font.Hinting(font.WeightLight),
	})

	messageInnerMplusNormalFont := truetype.NewFace(ttMedium, &truetype.Options{
		Size:    16,
		DPI:     72,
		Hinting: font.Hinting(font.WeightLight),
	})

	menu.fontFaces[NAME] = nameMplusNormalFont
	menu.fontFaces[MENU] = messageMplusNormalFont
	menu.fontFaces[MENU_INNER] = messageInnerMplusNormalFont
}

func (menu *StateMenu) OnDestroy() {}

func (menu *StateMenu) Activate() {
	menu.eventMgr.AddCallback(schema.Menu, "KeyR", func(ed *event.EventDetail) {
		menu.isMenuItemDown[RESUME] = true
	})
	menu.eventMgr.AddCallback(schema.Menu, "KeyRUp", func(ed *event.EventDetail) {
		menu.stateMgr.SwitchTo(schema.Game)
		menu.isMenuItemDown[RESUME] = false
	})
	menu.eventMgr.AddCallback(schema.Menu, "KeyQ", func(ed *event.EventDetail) {
		menu.isMenuItemDown[QUIT] = true
	})
	menu.eventMgr.AddCallback(schema.Menu, "KeyQUp", func(ed *event.EventDetail) {
		menu.stateMgr.SwitchTo(schema.Intro)
		menu.isMenuItemDown[QUIT] = false
	})
	menu.eventMgr.AddCallback(schema.Menu, "KeyC", func(ed *event.EventDetail) {
		menu.isMenuItemDown[CONTROL] = true
		menu.ToggleShowControl()
	})
	menu.eventMgr.AddCallback(schema.Menu, "KeyCUp", func(ed *event.EventDetail) {
		menu.isMenuItemDown[CONTROL] = false
	})
}

func (menu *StateMenu) Deactivate() {
	menu.eventMgr.RemoveCallback(schema.Intro, "KeyR")
	menu.eventMgr.RemoveCallback(schema.Game, "KeyQ")
	menu.eventMgr.RemoveCallback(schema.Game, "KeyC")
	menu.eventMgr.RemoveCallback(schema.Intro, "KeyRUp")
	menu.eventMgr.RemoveCallback(schema.Game, "KeyQUp")
	menu.eventMgr.RemoveCallback(schema.Game, "KeyCUp")
}

func (menu *StateMenu) Update(elapsed time.Duration) {}

func (menu *StateMenu) Render(screen *ebiten.Image) {
	x, y := ebiten.WindowSize()

	for i := 0; i < int(x); i++ {
		if i > int((x-menuWidth)/2) && i < x-int((x-menuWidth)/2) {
			vector.DrawFilledRect(screen, float32(i)+3, float32(y-menuHeight)/2-3, 1, float32(menuHeight), color.NRGBA{0xc4, 0xdd, 0xff, 0x77}, false)
			vector.DrawFilledRect(screen, float32(i)-3, float32(y-menuHeight)/2+3, 1, float32(menuHeight), color.NRGBA{0xc4, 0xdd, 0xff, 0x77}, false)
		}
	}

	textColor := color.RGBA{0xff, 0xff, 0xff, 0xff}
	chosenTextColor := color.RGBA{0xa2, 0xdb, 0xfa, 0xff}

	menu.RenderGameName(screen)

	if menu.isMenuItemDown[RESUME] {
		menu.RenderResumeItem(screen, chosenTextColor)
	} else {
		menu.RenderResumeItem(screen, textColor)
	}

	if menu.isMenuItemDown[QUIT] {
		menu.RenderQuitItem(screen, chosenTextColor)
	} else {
		menu.RenderQuitItem(screen, textColor)
	}

	if menu.isMenuItemDown[CONTROL] {
		menu.RenderControlItem(screen, chosenTextColor)
	} else {
		menu.RenderControlItem(screen, textColor)
	}
}

func (menu *StateMenu) RenderGameName(screen *ebiten.Image) {
	x, y := ebiten.WindowSize()
	nameLabel := "Go\nWorm"
	nameBound := text.BoundString(menu.fontFaces[NAME], nameLabel)
	namePosY := (float64(y) - (float64(nameBound.Max.Y)-float64(nameBound.Min.Y))/4) / 2
	text.Draw(screen, nameLabel, menu.fontFaces[NAME], int(x/2)+57, int(namePosY), color.RGBA{0xc4, 0xdd, 0xff, 0xff})
	text.Draw(screen, nameLabel, menu.fontFaces[NAME], int(x/2)+60, int(namePosY), color.RGBA{0x65, 0x28, 0xf7, 0xff})
}

func (menu *StateMenu) RenderResumeItem(screen *ebiten.Image, color color.RGBA) {
	x, y := ebiten.WindowSize()
	label := "Resume (R)"
	text.Draw(screen, label, menu.fontFaces[MENU], int((x-menuWidth)/2)+25, int((y-menuHeight)/2)+40, color)
}

func (menu *StateMenu) RenderQuitItem(screen *ebiten.Image, color color.RGBA) {
	x, y := ebiten.WindowSize()
	label := "Quit (Q)"
	text.Draw(screen, label, menu.fontFaces[MENU], int((x-menuWidth)/2)+25, int((y-menuHeight)/2)+75, color)
}

func (menu *StateMenu) RenderControlItem(screen *ebiten.Image, textColor color.RGBA) {
	x, y := ebiten.WindowSize()
	label := "Control (C)"
	detail := "• Direction: Comma (,), Period (.)\n\n• Fire (hold to charge): X\n\n• Next player: N"

	if menu.showControl {
		text.Draw(screen, label, menu.fontFaces[MENU], int((x-menuWidth)/2)+25, int((y-menuHeight)/2)+110, textColor)
		text.Draw(screen, detail, menu.fontFaces[MENU_INNER], int((x-menuWidth)/2)+40, int((y-menuHeight)/2)+140, color.RGBA{0xff, 0xff, 0xff, 0xff})
	} else {
		text.Draw(screen, label, menu.fontFaces[MENU], int((x-menuWidth)/2)+25, int((y-menuHeight)/2)+110, textColor)
	}
}

func (menu *StateMenu) IsTransparent() bool {
	return true
}

func (menu *StateMenu) IsTranscendent() bool {
	return false
}

func (menu *StateMenu) ToggleShowControl() {
	menu.showControl = !menu.showControl
}
