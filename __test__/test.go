package test

import (
	"main/callstack"
	eventlist "main/constants/event-list"
	"main/model"
)

func Test() error {
	callstack.PushCallStack(model.Event{
		Key:     eventlist.PURE_REMOVE_BAT,
		KeyWord: "test",
	})
	callstack.PushCallStack(model.Event{
		Key:     eventlist.RUNBAT,
		KeyWord: "test",
	})
	callstack.PushCallStack(model.Event{
		Key:     eventlist.BAT,
		KeyWord: "test",
		Payload: "echo 67",
	})

	return nil
}
