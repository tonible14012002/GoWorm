package state

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/event"
	"github.com/tonible14012002/go_game/engine/schema"
	"golang.org/x/exp/slices"
)

type StateManager struct {
	states             StateStack
	toRemoveStateTypes []schema.StateType
	factory            StateFactory
	currentState       schema.StateType
	eventMgr           *event.EventManager
}

func (stateMgr *StateManager) RegisterEventManager(eventMgr *event.EventManager) {
	stateMgr.eventMgr = eventMgr

}

func (stateMgr *StateManager) Setup() {
	totalStates := len(schema.GetAllGameStateTypes())
	stateMgr.factory = make(StateFactory)
	stateMgr.states = make(StateStack, 0, totalStates)
	stateMgr.toRemoveStateTypes = make([]schema.StateType, 0, totalStates)
	if stateMgr.eventMgr == nil {
		log.Fatal("state manager must initialize with event manager")
	}
}

func (stateMgr *StateManager) GetAllStateFactories() StateFactory {
	return stateMgr.factory
}

func (stateMgr *StateManager) RegisterState(stateType schema.StateType, generator StateGenerator) {
	stateMgr.factory[stateType] = generator
}

func (stateMgr *StateManager) createState(stateType schema.StateType) (*StateInfo, error) {
	// Create New Registered State
	if stateFactory, exist := stateMgr.factory[stateType]; exist {
		newStateInfoPtr := &StateInfo{
			Statetype: stateType,
			GameState: stateFactory(),
		}
		stateMgr.states = append(stateMgr.states, newStateInfoPtr)
		newStateInfoPtr.GameState.OnCreate(stateMgr, stateMgr.eventMgr)
		return newStateInfoPtr, nil
	}
	return nil, errors.New("no State type match when creatstate")
}

func (stateMgr *StateManager) SwitchTo(sType schema.StateType) {
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
		fmt.Println("Found state")
		fmt.Println("Activate", foundState.Statetype)
		foundState.GameState.Activate()
		stateMgr.SetCurrentState(sType)
		return
	}

	fmt.Println("Not Found state, creating...")
	if len(stateMgr.states) != 0 {
		lastState := stateMgr.states[len(stateMgr.states)-1]
		lastState.GameState.Deactivate()
		fmt.Println("Dactivate", lastState.Statetype)
	}

	newState, err := stateMgr.createState(sType)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Activate", sType)
	newState.GameState.Activate()
	// State not found in stack
	stateMgr.SetCurrentState(sType)
	fmt.Println(stateMgr.states)
}

func (stateMgr *StateManager) GetCurrentState() schema.StateType { return stateMgr.currentState }
func (stateMgr *StateManager) SetCurrentState(sType schema.StateType) {
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
	stateMgr.toRemoveStateTypes = make([]schema.StateType, 0, len(schema.GetAllGameStateTypes()))
}

func (stateMgr *StateManager) Remove(stateType schema.StateType) {
	stateMgr.toRemoveStateTypes = append(stateMgr.toRemoveStateTypes, stateType)
}

func (stateMgr *StateManager) HasState(stateType schema.StateType) bool {
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

func (stateMgr *StateManager) Update(elapsed time.Duration) {
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

func (stateMgr *StateManager) Render(screen *ebiten.Image) {
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
			states[i].GameState.Render(screen)
		}
	} else {
		states[i].GameState.Render(screen)
	}
}
