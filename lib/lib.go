package lib

import (
	"main/callstack/event"
	"main/callstack/manual"
	"main/constants"
	eventlist "main/constants/event-list"
	"main/lib/death"
	"main/lib/handler"
	"main/model"
	"main/terminal"
)

func Init() {
	if err := constants.Init_root(); err != nil {
		terminal.Println("Ошибка инициализации")
		return
	}
	terminal.RenderBaner()
	manual.Manual()
}

func Event(e model.Event) error {
	return handler.Dispatch(e)
}

func HandleEvent(e model.Event) error {
	if err := Event(e); err != nil {
		event.Response("Ошибка: ", err)
		if logErr := death.Logger(e); logErr != nil {
			event.Response("Ошибка лога: ", logErr)
		}
		return err
	}

	if e.Key != eventlist.RESPONSE {
		event.Success()
	}

	if err := death.Logger(e); err != nil {
		event.Response("Ошибка лога: ", err)
		return err
	}
	return nil
}
