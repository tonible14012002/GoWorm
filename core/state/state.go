package state

type StateManager struct {
	states             StateStack
	toRemoveStateTypes []StateType
	factory            StateFactory
}

func (stateManager *StateManager) Setup() {
	totalStates := len(GetAllGameStateTypes())
	stateManager.factory = make(StateFactory)
	stateManager.states = make(StateStack, 0, totalStates)
	stateManager.toRemoveStateTypes = make([]StateType, 0, totalStates)
}

func (stateManager *StateManager) RegisterState(stateType StateType, generator StateGenerator) {
	stateManager.factory[stateType] = generator
}

func (stateManager *StateManager) createState(stateType StateType) {

}

func (stateManager *StateManager) SwitchTo(StateType StateType) {

}

func (stateManager *StateManager) processRemoveRequest() {
	// This method must call at the end of Global Update
	toRemoveIndices := make([]int, 0, len(stateManager.toRemoveStateTypes))
	for _, toRemoveStateType := range stateManager.toRemoveStateTypes {
		for index, stateInfo := range stateManager.states {
			if stateInfo.Statetype == toRemoveStateType {
				toRemoveIndices = append(toRemoveIndices, index)
			}
		}
	}
	// Remove
	// for _, index := range toRemoveIndices {

	// }
}

func (stateManager *StateManager) Remove(stateType StateType) {
	stateManager.toRemoveStateTypes = append(stateManager.toRemoveStateTypes, stateType)
}

func (stateManager *StateManager) HasState(stateType StateType) bool {
	// TRUE: stateType contained in stateStack and not wont be remove
	for _, stateInfo := range stateManager.states {
		if stateInfo.Statetype == stateType {
			for _, toRemove := range stateManager.toRemoveStateType {
				if toRemove == stateType {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (stateManager *StateManager) Update(elapsed float64) {
	if len(stateManager.states) == 0 {
		return
	}
	// render lower layer if current layer is transparent
	states := stateManager.states
	i := len(states) - 1
	if states[i].GameState.IsTranscendent() && i > 0 {
		for i != 0 {
			if !states[i].GameState.IsTranscendent() {
				break
			}
			i--
		}
		for ; i < len(states); i++ {
			states[i].GameState.Update(elapsed)
		}
	} else {
		states[i].GameState.Update(elapsed)
	}
}

func (stateManager *StateManager) Render() {
	if len(stateManager.states) == 0 {
		return
	}
	// render lower layer if current layer is transparent
	states := stateManager.states
	i := len(states) - 1
	if states[i].GameState.IsTransparent() && i > 0 {
		for i != 0 {
			if !states[i].GameState.IsTransparent() {
				break
			}
			i--
		}
		for ; i < len(states); i++ {
			states[i].GameState.Render()
		}
	} else {
		states[i].GameState.Render()
	}
}
