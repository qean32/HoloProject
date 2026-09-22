package lib

import (
	"bufio"
	"main/constants"
	"main/lib/handler"
	"main/lib/low"
	"main/lib/manual"
	"main/model"
	"main/terminal"
	"main/terminal/list"
	"os"
)

var READER = bufio.NewReader(os.Stdin)

func INIT() {
	terminal.OutputASCII_CENTER(constants.BinaryPROJECT_INIT, " ")
	terminal.OutputASCII_CENTER(constants.PROJECT_INIT, " ")
	constants.INIT_ROOT()
	list.List(handler.Menu)
}

func Event(e model.Event) {
	defer manual.Manual()
	function := handler.MAP[e.Key]

	if function != nil {
		function(e)
		low.LOG(e)
	} else {
		terminal.Outputln(constants.UNDEFINED_COMMAND)
	}
}
