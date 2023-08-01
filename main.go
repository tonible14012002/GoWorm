package main

import "fmt"

// type Game struct {
// 	ebiten.Game
// }
// func (game *Game) Update() error {
// 	return nil
// }

// func (game *Game) Draw(screen *ebiten.Image) {

// }

// func (game *Game) Layout(outerWidth, outerHeight int) (screenWidth, screenHeight int) {
// 	return 640, 320
// }

func main() {
	a := map[string]int{
		"1": 1,
	}
	Print("a", a)

}

func Print(name string, a interface{}) {
	fmt.Printf("address of %s = %p\n", name, a)
}
