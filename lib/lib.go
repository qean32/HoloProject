package lib

import (
	"bufio"
	"os"

	"main/constants"
	"main/constants/response"
	"main/lib/handler"
	"main/lib/low"
	"main/lib/manual"
	"main/model"
	"main/terminal"
)

var READER = bufio.NewReader(os.Stdin)

func INIT() {
	terminal.PrintASCIICenter(constants.BinaryPROJECT_INIT, " ")
	terminal.PrintASCIICenter(constants.PROJECT_INIT, " ")
	constants.INIT_ROOT()
	low.SETDATA()
	manual.Manual()
}

func Event(e model.Event) {
	function := handler.MAP[e.Key]

	if function == nil {
		Event(response.UndefinedCommand())
		return
	}
	function(e)
	low.LOG(e)
}
