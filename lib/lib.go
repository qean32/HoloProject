package lib

import (
	"bufio"
	"os"

	"main/callstack/event"
	"main/callstack/manual"
	"main/constants"
	"main/lib/death"
	"main/lib/handler"
	"main/model"
	"main/terminal"
)

var READER = bufio.NewReader(os.Stdin)

func INIT() {
	terminal.PrintASCIICenter(constants.BinaryPROJECT_INIT, " ")
	terminal.PrintASCIICenter(constants.PROJECT_INIT, " ")
	constants.INIT_ROOT()
	death.SETDATA()
	manual.Manual()
}

func Event(e model.Event) {
	function := handler.MAP[e.Key]

	if function == nil {
		event.UndefinedCommand()
		return
	}
	function(e)
	death.Logger(e)
}
