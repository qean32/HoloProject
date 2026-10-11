package test

import (
	"main/callstack"
	eventlist "main/constants/event-list"
	"main/model"
)

func Test() {
	callstack.PushCallStack(model.Event{
		Key:     eventlist.RRBAT,
		KeyWord: "test",
	})
	callstack.PushCallStack(model.Event{
		Key:     eventlist.RUNBAT,
		KeyWord: "test",
	})
	callstack.PushCallStack(model.Event{
		Key:     eventlist.BAT,
		KeyWord: "test",
		Payload: "start brave",
	})
}
