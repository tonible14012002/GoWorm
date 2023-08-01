package event

type Event struct {
	// key code
}

type EventDetail struct {
	// TODO:
	// data structure for storing the event detail
	// shared for the code use this event
}

type Binding struct {
	// TODO: Binding detail for
	// implement binding key event
}

type BindingMap map[string]Binding

type EventManager struct {
}

func (eventManager *EventManager) LoadBinding() {
	//TODO:  Load Input binding from config.txt
}

func (eventManager *EventManager) AddBinding() {

}

func (eventManager *EventManager) RemoveBinding() {

}

func (EventManager *EventManager) AddCallback() {

}

func (EventManager *EventManager) RemoveCallback() {

}

func (eventManager *EventManager) Update() {
	// TODO: Implement handle Event
}
