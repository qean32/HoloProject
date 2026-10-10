package event

import (
	"main/callstack"
	"main/constants/event-list"
	"main/model"
)

func UndefinedCommand() {
	callstack.PushCallStack(model.Event{
		Key:     eventlist.RESPONSE,
		Payload: "Неизвестная команда",
	})
}

func SyntaxError() {
	callstack.PushCallStack(model.Event{
		Key:     eventlist.RESPONSE,
		Payload: "Syntax error",
	})
}

func Success() {
	callstack.PushCallStack(model.Event{
		Key:     eventlist.RESPONSE,
		Payload: "/ᐠ-˕-マᶻ",
	})
}

func UnInit() {
	callstack.PushCallStack(model.Event{
		Key:     eventlist.RESPONSE,
		Payload: "Отсутствует инициализация! -init",
	})
}
