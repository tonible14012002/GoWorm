package event

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/schema"
)

type EventType int

const (
	KeyDown EventType = iota
	KeyUp
	KeyPressed
	MouseDown
	MouseUp
	MousePressed
	MouseWheel
	MouseEnter
	MouseLeave
	Resized
	GainFocus
	LostFocus
	Closed
)

var KeyEvents = [3]EventType{KeyDown, KeyUp, KeyPressed}
var MouseEvents = [3]EventType{MouseDown, MouseUp, MousePressed}
var MouseMoveEvents = [2]EventType{MouseEnter, MouseLeave}

type Event struct {
	// Hold the Event Type and keyCode of Mouse/KeyBoard
	eType     EventType
	keyCode   ebiten.Key
	mouseCode ebiten.MouseButton
}

type EventDetail struct {
	// Hold all the Event Detail for sharing
	Name        string
	TextEntered rune //
	KeyCode     int  //
	MouseCode   int  //

	MouseWheelX float64
	MouseWheelY float64
	MouseX      int
	MouseY      int
}
type Binding struct {
	// Hold binding info
	name           string
	events         []Event
	happeningCount int
	detail         EventDetail
}

type CallBack func(*EventDetail)

func (b *Binding) BindEvent(e Event) {
	b.events = append(b.events, e)
}

type KeyEventChecker func(e ebiten.Key) bool

type MouseEventChecker func(e ebiten.MouseButton) bool

type CallBackDict map[string]CallBack
type StateCallBackDict map[schema.StateType]CallBackDict
type BindingDict map[string]*Binding
