package event

import (
	"main/callstack"
	eventlist "main/constants/event-list"
	"main/model"
)

func Success() {
	callstack.Push(model.Event{
		Key:     eventlist.RESPONSE,
		Payload: "/ᐠ-˕-マ Успех!",
	})
}
