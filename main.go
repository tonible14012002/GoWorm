package main

import (
	"log"

	"github.com/tonible14012002/go_game/engine/game"
)

func main() {
	game := game.Game{}
	game.Setup()
	if err := game.Start(); err != nil {
		log.Fatal(err)
		return
	}
}
