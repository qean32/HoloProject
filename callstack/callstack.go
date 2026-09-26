package callstack

import (
	"main/model"
	"main/terminal"
	"sync"
)

var (
	mu      sync.Mutex
	queue   []model.Event
	changed = make(chan struct{}, 1)
)

func PushCallStack(event model.Event) {
	mu.Lock()
	queue = append(queue, event)
	mu.Unlock()

	select {
	case changed <- struct{}{}:
	default:
	}
}

func Shift() (model.Event, bool) {
	mu.Lock()
	defer mu.Unlock()

	if len(queue) == 0 {
		return model.Event{}, false
	}

	event := queue[0]
	queue[0] = model.Event{}
	queue = queue[1:]
	return event, true
}

func Pop() (model.Event, bool) {
	mu.Lock()
	defer mu.Unlock()

	if len(queue) == 0 {
		return model.Event{}, false
	}

	terminal.PrintTechInfo(queue)
	last := len(queue) - 1
	event := queue[last]
	queue[last] = model.Event{} // обнуляем ссылку для GC
	queue = queue[:last]
	return event, true
}

func Changed() <-chan struct{} {
	return changed
}
