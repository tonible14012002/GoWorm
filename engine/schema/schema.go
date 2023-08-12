package schema

type StateType int

const Global StateType = -1
const (
	None StateType = iota
	Intro
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
