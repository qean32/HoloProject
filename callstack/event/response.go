package event

import (
	"main/callstack"
	"main/constants/literals"
	"main/model"
)

func UndefinedCommand() {
	callstack.PushCallStack(model.Event{
		Key:     literals.EventList.RESPONSE,
		Payload: literals.Response.UndefinedCommand,
	})
}

func SyntaxError() {
	callstack.PushCallStack(model.Event{
		Key:     literals.EventList.RESPONSE,
		Payload: literals.Response.SyntaxError,
	})
}

func Success() {
	callstack.PushCallStack(model.Event{
		Key:     literals.EventList.RESPONSE,
		Payload: literals.Response.Success,
	})
}
