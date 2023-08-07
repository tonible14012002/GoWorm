package state

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/exp/slices"
)

type StateManager struct {
	states             StateStack
	toRemoveStateTypes []StateType
	factory            StateFactory
	currentState       StateType
}

func (stateMgr *StateManager) Setup() {
	totalStates := len(GetAllGameStateTypes())
	stateMgr.factory = make(StateFactory)
	stateMgr.states = make(StateStack, 0, totalStates)
	stateMgr.toRemoveStateTypes = make([]StateType, 0, totalStates)
}

func (stateMgr *StateManager) GetAllStateFactories() StateFactory {
	return stateMgr.factory
}

func (stateMgr *StateManager) RegisterState(stateType StateType, generator StateGenerator) {
	stateMgr.factory[stateType] = generator
	fmt.Println(stateMgr.factory)
}

func (stateMgr *StateManager) createState(stateType StateType) (*StateInfo, error) {
	// Create New Registered State
	if stateFactory, exist := stateMgr.factory[stateType]; exist {
		newStateInfoPtr := &StateInfo{
			Statetype: stateType,
			GameState: stateFactory(),
		}
		stateMgr.states = append(stateMgr.states, newStateInfoPtr)
		newStateInfoPtr.GameState.OnCreate()
		return newStateInfoPtr, nil
	}
	return nil, errors.New("no State type match when creatstate")
}

func (stateMgr *StateManager) SwitchTo(sType StateType) {
	// Active new State
	// If not in stack, Create new state
	if sType == stateMgr.currentState {
		return
	}

	foundIndex := -1
	var foundState *StateInfo = nil

	for i := range stateMgr.states {
		if stateMgr.states[i].Statetype == sType {
			lastState := stateMgr.states[len(stateMgr.states)-1]
			lastState.GameState.Deactivate()
			foundIndex = i
			foundState = stateMgr.states[i]
			break
		}
	}
	if foundIndex != -1 {
		stateMgr.states = append(stateMgr.states[0:foundIndex], stateMgr.states[foundIndex+1:]...)
		stateMgr.states = append(stateMgr.states, foundState)
		return
	}
	if len(stateMgr.states) != 0 {
		lastState := stateMgr.states[len(stateMgr.states)-1]
		lastState.GameState.Deactivate()
	}

	newState, err := stateMgr.createState(sType)
	if err != nil {
		log.Fatal(err)
	}
	newState.GameState.Activate()
	// State not found in stack
	stateMgr.SetCurrentState(sType)
}

func (stateMgr *StateManager) GetCurrentState() StateType { return stateMgr.currentState }
func (stateMgr *StateManager) SetCurrentState(sType StateType) {
	stateMgr.currentState = sType
}

func (stateMgr *StateManager) LateUpdate() {
	stateMgr.processRemoveRequest()
}

func (stateMgr *StateManager) processRemoveRequest() {
	// Destroy all to remove state
	remainStatesSize := len(stateMgr.states) - len(stateMgr.toRemoveStateTypes)
	remainStates := make(StateStack, 0, remainStatesSize)

	for _, stateInfoPtr := range stateMgr.states {
		if !slices.Contains(stateMgr.toRemoveStateTypes, stateInfoPtr.Statetype) {
			remainStates = append(remainStates, stateInfoPtr)
		} else {
			stateInfoPtr.GameState.OnDestroy()
		}
	}
	stateMgr.states = remainStates
	stateMgr.toRemoveStateTypes = make([]StateType, 0, len(GetAllGameStateTypes()))
}

func (stateMgr *StateManager) Remove(stateType StateType) {
	stateMgr.toRemoveStateTypes = append(stateMgr.toRemoveStateTypes, stateType)
}

func (stateMgr *StateManager) HasState(stateType StateType) bool {
	// TRUE: stateType contained in stateStack and not wont be remove
	for _, stateInfoPtr := range stateMgr.states {
		if stateInfoPtr.Statetype == stateType {
			for _, toRemove := range stateMgr.toRemoveStateTypes {
				if toRemove == stateType {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (stateMgr *StateManager) Update(elapsed float64) {
	if len(stateMgr.states) == 0 {
		return
	}
	// render lower layer if current layer is transparent
	stateInfoPtrs := stateMgr.states
	i := len(stateInfoPtrs) - 1
	if stateInfoPtrs[i].GameState.IsTranscendent() && i > 0 {
		for i != 0 {
			if !stateInfoPtrs[i].GameState.IsTranscendent() {
				break
			}
			i--
		}
		for ; i < len(stateInfoPtrs); i++ {
			stateInfoPtrs[i].GameState.Update(elapsed)
		}
	} else {
		stateInfoPtrs[i].GameState.Update(elapsed)
	}
}

func (stateMgr *StateManager) Render() {
	if len(stateMgr.states) == 0 {
		return
	}
	// render lower layer if current layer is transparent
	states := stateMgr.states
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
