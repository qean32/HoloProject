package event

import (
	"main/callstack"
	eventlist "main/constants/event-list"
	"main/model"
)

func UndefinedCommand() {
	callstack.Push(model.Event{
		Key:     eventlist.RESPONSE,
		Payload: "Неизвестная команда",
	})
}

func SyntaxError() {
	callstack.Push(model.Event{
		Key:     eventlist.RESPONSE,
		Payload: "Syntax error",
	})
}

func Success() {
	callstack.Push(model.Event{
		Key:     eventlist.RESPONSE,
		Payload: "/ᐠ-˕-マᶻ",
	})
}

func UnInit() {
	callstack.Push(model.Event{
		Key:     eventlist.RESPONSE,
		Payload: "Отсутствует инициализация! -init",
	})
}
