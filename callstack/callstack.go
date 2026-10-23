package callstack

import (
	"sync"

	"main/model"
)

var (
	mu    sync.Mutex
	queue []model.Event
)

func Push(event model.Event) {
	mu.Lock()
	queue = append(queue, event)
	mu.Unlock()
}

func PushMany(events ...model.Event) {
	if len(events) == 0 {
		return
	}

	mu.Lock()
	queue = append(queue, events...)
	mu.Unlock()
}

func Pop() (model.Event, bool) {
	mu.Lock()
	defer mu.Unlock()

	if len(queue) == 0 {
		return model.Event{}, false
	}

	last := len(queue) - 1
	event := queue[last]
	queue[last] = model.Event{}
	queue = queue[:last]
	return event, true
}
