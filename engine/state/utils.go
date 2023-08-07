package state

func GetAllGameStateTypes() []StateType {
	return []StateType{
		Intro,
		Menu,
		Game,
		Ending,
	}
}
