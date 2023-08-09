package schema

type StateType int

const (
	Global StateType = -1
	Intro  StateType = iota
	Menu
	Game
	Ending
)

func GetAllGameStateTypes() []StateType {
	return []StateType{
		Intro,
		Menu,
		Game,
		Ending,
	}
}
