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
	"slices"
)

var succesList = []string{
	eventlist.BAT,
	eventlist.PURE_REMOVE_BAT,
	eventlist.TEST,
	eventlist.CLEARLOG,
	eventlist.CRYPTO,
	eventlist.DROP,
	eventlist.INIT,
	eventlist.RUNBAT,
}

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
		event.Response(err)
		if logErr := death.Logger(e); logErr != nil {
			event.Response(logErr)
		}
		return err
	}

	if slices.Contains(succesList, e.Key) {
		event.Success()
	}

	if err := death.Logger(e); err != nil {
		event.Response(err)
		return err
	}
	return nil
}
