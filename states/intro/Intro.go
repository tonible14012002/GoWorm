package intro

import "fmt"

type StateIntro struct {
	isTransparent  bool
	isTranscendent bool
	name           string
	isActivated    bool
	isUpdated      bool
}

func (intro *StateIntro) OnCreate() {
	fmt.Println("OnCreating...")
	intro.name = "Creating"
	intro.isActivated = false
	intro.isUpdated = false
}

func (intro *StateIntro) OnDestroy() {
	fmt.Println("Destroying...")
}

func (intro *StateIntro) Activate() {
	// Begin
	fmt.Println("Activate...")
}

func (intro *StateIntro) Deactivate() {
	fmt.Println("Deactivate...")
}

func (intro *StateIntro) Update(elapsed float64) {
	if !intro.isUpdated {
		fmt.Println("update state")
	}
	intro.isUpdated = true
}

func (intro *StateIntro) Render() {

}

func (intro *StateIntro) SetTransparent(isTransparent bool) {
	intro.isTransparent = isTransparent
}

func (intro *StateIntro) IsTransparent() bool {
	return intro.isTransparent
}

func (intro *StateIntro) SetTranscendent(isTranscendent bool) {
	intro.isTranscendent = isTranscendent
}

func (intro *StateIntro) IsTranscendent() bool {
	return intro.isTranscendent
}
