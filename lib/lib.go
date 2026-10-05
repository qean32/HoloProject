package lib

import (
	"bufio"
	"os"

	"main/callstack/event"
	"main/callstack/manual"
	"main/constants"
	eventlist "main/constants/event-list"
	"main/lib/handler"
	"main/model"
	"main/terminal"
)

var READER = bufio.NewReader(os.Stdin)

func INIT() {
	constants.INIT_ROOT()
	terminal.RenderBaner()
	manual.Manual()
}

func Event(event_ model.Event) {
	if event_.Key != eventlist.RESPONSE {
		defer event.Success()
	}

	if err := handler.Dispatch(event_); err != nil {
		event.Response("Ошибка: ", err)
		// if logErr := death.Logger(event_); logErr != nil {
		// 	event.Response("Ошибка лога: ", logErr)
		// }
		return
	}

	// if err := death.Logger(event_); err != nil {
	// 	event.Response("Ошибка лога: ", err)
	// }
}
