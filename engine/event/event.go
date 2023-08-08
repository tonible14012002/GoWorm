package event

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonible14012002/go_game/engine/state"
)

type CallBackDict map[string]CallBack
type StateCallBackDict map[state.StateType]CallBackDict
type BindingDict map[string]*Binding

type EventManager struct {
	bindings  BindingDict
	callbacks StateCallBackDict
}

func (eManager *EventManager) Setup() {
	eManager.callbacks = make(StateCallBackDict)
	eManager.bindings = make(BindingDict)

	states := state.GetAllGameStateTypes()
	for _, stateType := range states {
		emptyDict := make(CallBackDict)
		eManager.callbacks[stateType] = emptyDict
	}

	err := eManager.loadBinding()
	if err != nil {
		log.Fatal("Cannot read key config file.")
	}
}

func (eManager *EventManager) AddCallback(gState state.StateType, name string, f CallBack) {
	if _, exist := eManager.callbacks[gState]; !exist {
		eManager.callbacks[gState] = make(CallBackDict)
	}
	eManager.callbacks[gState][name] = f
}

func (eManager *EventManager) RemoveCallback(gState state.StateType, name string) bool {
	if _, exist := eManager.callbacks[gState]; !exist {
		return false
	}
	delete(eManager.callbacks[gState], name)
	return true
}

func (eManager *EventManager) addBinding(bind *Binding) bool {
	if _, exist := eManager.bindings[bind.name]; exist {
		return false
	}
	eManager.bindings[bind.name] = bind
	return true
}

func (eManager *EventManager) loadBinding() error {
	file, err := os.Open("configuration.cfg")
	if err != nil {
		return err
	}
	fileSCanner := bufio.NewScanner(file)
	configs := make([]string, 0, 10)
	for fileSCanner.Scan() {
		configs = append(configs, fileSCanner.Text())
	}

	for i := range configs {
		keyConfig := &configs[i]
		sections := strings.Split(*keyConfig, " ")
		if len(sections) < 2 {
			return errors.New("incorrect Format config file")
		}
		bindName := sections[0]
		events := sections[1:]

		newBinding := Binding{name: bindName}

		for _, event := range events {
			parts := strings.Split(event, ":")
			if len(parts) != 2 {
				return fmt.Errorf("incorrect Event key format in line %d", i)
			}
			eType, eErr := strconv.Atoi(parts[0])
			code, cErr := strconv.Atoi(parts[1])
			if eErr != nil || cErr != nil {
				return errors.New("inccorect Keycode/EventType, must be an integer")
			}
			newBinding.events = append(newBinding.events, Event{eType: EventType(eType), keyCode: ebiten.Key(code), mouseCode: ebiten.MouseButton(code)})
		}
		eManager.addBinding(&newBinding)
	}
	file.Close()
	return nil
}

func (eManager *EventManager) Update(currentState state.StateType) {
	for _, binding := range eManager.bindings {
		for _, event := range binding.events {
			if isKeyEvent(event.eType) {
				keyChecker := getKeyEventChecker(event.eType)
				active := keyChecker(event.keyCode)
				if active {
					binding.happeningCount++
					if binding.detail.KeyCode != -1 {
						binding.detail.KeyCode = int(event.keyCode)
					}
				}
			} else if isMouseEvent(event.eType) {
				mouseChecker := getMouseEventChecker(event.eType)
				active := mouseChecker(event.mouseCode)
				if active {
					binding.happeningCount++
					mouseX, mouseY := ebiten.CursorPosition()
					binding.detail.MouseX = mouseX
					binding.detail.MouseY = mouseY
					if binding.detail.MouseCode != -1 {
						binding.detail.MouseCode = int(event.mouseCode)
					}
				}
			} else {
				switch event.eType {
				case MouseWheel:
					wheelX, wheelY := ebiten.Wheel()
					if wheelX != 0 || wheelY != 0 {
						binding.happeningCount++
						binding.detail.MouseWheelX = wheelX
						binding.detail.MouseWheelY = wheelY
					}
				case GainFocus:
				case Resized:
					// Updateing ...
				}
			}
			if binding.happeningCount == len(binding.events) {
				currentStateCallback, curExist := eManager.callbacks[currentState][binding.name]
				globalStateCallback, gloExist := eManager.callbacks[state.Global][binding.name]
				if curExist {
					currentStateCallback(&binding.detail)
				}
				if gloExist {
					globalStateCallback(&binding.detail)
				}
			}
			binding.happeningCount = 0
		}
	}
}
