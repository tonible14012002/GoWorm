package main

import (
	"log"

	"github.com/tonible14012002/go_game/engine/gamemanager"
)

func main() {
	gameManager := gamemanager.Game{}
	gameManager.Setup()
	if err := gameManager.Start(); err != nil {
		log.Fatal(err)
		return
	}
}
