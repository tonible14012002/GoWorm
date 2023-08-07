package menu

import "fmt"

type StateGame struct {
	isTransparent  bool
	isTranscendent bool
	name           string
	isActivated    bool
}

func (intro *StateGame) OnCreate() {
	intro.name = "Creating"
	intro.isActivated = true
	fmt.Println("Creating")
}

func (intro *StateGame) OnDestroy() {
	fmt.Println("Destroying")
}

func (intro *StateGame) Activate() {
	// Begin
	intro.isActivated = true
	fmt.Println("Activate")
}

func (intro *StateGame) Deactivate() {
	fmt.Println("Deactivate")
}

func (intro *StateGame) Update(elapsed float64) {
	if intro.isActivated {
		fmt.Println("Running Game")
	}
	intro.isActivated = false
}

func (intro *StateGame) Render() {

}

func (intro *StateGame) SetTransparent(isTransparent bool) {
	intro.isTransparent = isTransparent
}

func (intro *StateGame) IsTransparent() bool {
	return intro.isTransparent
}

func (intro *StateGame) SetTranscendent(isTranscendent bool) {
	intro.isTranscendent = isTranscendent
}

func (intro *StateGame) IsTranscendent() bool {
	return intro.isTranscendent
}
