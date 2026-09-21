package callstack

import (
	"sync"

	"main/model"
)

var (
	mu        sync.Mutex
	callstack []model.Event
	changed   = make(chan struct{}, 1)
	wasEmpty  = true
)

func PushCallStack(event model.Event) {
	mu.Lock()
	callstack = append(callstack, event)
	shouldNotify := wasEmpty
	wasEmpty = false
	mu.Unlock()

	if shouldNotify {
		select {
		case changed <- struct{}{}:
		default:
		}
	}
}

func GetTask() (model.Event, bool) {
	mu.Lock()
	defer mu.Unlock()

	if len(callstack) == 0 {
		wasEmpty = true
		return model.Event{}, false
	}

	event := callstack[0]
	callstack[0] = model.Event{}
	callstack = callstack[1:]
	return event, true
}

func Changed() <-chan struct{} {
	return changed
}
