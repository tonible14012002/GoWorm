package event

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func IsRepeatingKeyPressed(k ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)

	d := inpututil.KeyPressDuration(k)
	// fmt.Println("d", d)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func isKeyEvent(e EventType) bool {
	for _, val := range KeyEvents {
		if e == val {
			return true
		}
	}
	return false
}

func getKeyEventChecker(e EventType) KeyEventChecker {
	switch e {
	case KeyDown:
		return inpututil.IsKeyJustPressed
	case KeyUp:
		return inpututil.IsKeyJustReleased
	case KeyPressed:
		return ebiten.IsKeyPressed
	default:
		return func(e ebiten.Key) bool { return false }
	}
}

func isMouseEvent(e EventType) bool {
	for _, val := range MouseEvents {
		if e == val {
			return true
		}
	}
	return false
}

func getMouseEventChecker(e EventType) MouseEventChecker {
	switch e {
	case MouseDown:
		return inpututil.IsMouseButtonJustPressed
	case MouseUp:
		return inpututil.IsMouseButtonJustReleased
	case MousePressed:
		return ebiten.IsMouseButtonPressed
	default:
		return func(e ebiten.MouseButton) bool { return false }
	}
}

// func IsMouseMoveEvent(e EventType) bool {
// 	for _, val := range MouseEvents {
// 		if e == val {
// 			return true
// 		}
// 	}
// 	return false
// }
