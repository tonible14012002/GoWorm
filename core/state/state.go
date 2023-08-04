package state

type GameState int

const (
	Global GameState = -1
	Intro  GameState = iota
	Menu
	Game
	Ending
)

func GetAllGameStates() []GameState {
	return []GameState{
		Intro,
		Menu,
		Game,
		Ending,
	}
}
