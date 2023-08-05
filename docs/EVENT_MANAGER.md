# EVENT MANAGER

## Functions used
- ebiten
```go
IsKeyPressed(key ebiten.Key) bool
Wheel() xOff, yOff float64
// AppendInputChars(chars []runes) []runes
// IsFocused() bool
// IsFullScreen() bool
// IsWindowResizable() bool
// SetFullscreen(fullscreen bool)
// SetScreenTransparent(transparent bool)
```
- ebiten/inpututils
```go
// AppendPressedKeys(keys []ebiten.Key) []ebiten.Key
// AppendJustPressedKeys(keys []ebiten.Key) []ebiten.Key
// AppendJustReleasedKeys(keys []ebiten.Key) []ebiten.Key
// KeyPressDuration(key ebiten.Key) int
// MouseButtonPressDuration(btn ebiten.MouseButton) int
IsKeyJustPressed(key ebiten.Key) bool
IsKeyJustReleased(key ebiten.Key) bool
```
## Event Management
## Usage
Example
```go
// player.go
type Player struct {
    posX int
    // ...
}
func MoveRight (e *EventDetail) {
    posX++
}

// Window.go
type  Window struct {
    eManager EventManager
    currentState GameState
    p Player
}

func (w *Window) Setup () {
    // Initialize function for Window struct
    eManager.AddCallback(currentState, "Move_Right", p.MoveRight)
}

// configuration.cfg
Move_Right 0:1 // KeyDown:KeyA
```
Now the Function MoveRight will be trigger if user press the key `A`