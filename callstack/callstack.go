package callstack

import (
	"sync"

	"main/model"
)

var (
	mu        sync.Mutex
	callstack []model.Event

	changed = make(chan struct{}, 1)
)

func PushCallStack(event model.Event) {
	mu.Lock()
	callstack = append(callstack, event)
	mu.Unlock()

	notify()
}

func GetTask() (model.Event, bool) {
	mu.Lock()
	defer mu.Unlock()

	if len(callstack) == 0 {
		return model.Event{}, false
	}

	event := callstack[0]
	callstack[0] = model.Event{}
	callstack = callstack[1:]
	return event, true
}

func Snapshot() []model.Event {
	mu.Lock()
	defer mu.Unlock()
	return append([]model.Event(nil), callstack...)
}

func Changed() <-chan struct{} {
	return changed
}

func notify() {
	select {
	case changed <- struct{}{}:
	default:
	}
}
