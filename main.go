package main

import (
	"log"

	"github.com/tonible14012002/go_game/engine/gamemanager"
)

func main() {
	game := gamemanager.Game{}
	game.Setup()
	if err := game.Start(); err != nil {
		log.Fatal(err)
		return
	}
}
