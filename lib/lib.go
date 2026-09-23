package lib

import (
	"bufio"
	"main/constants"
	"main/constants/response"
	"main/lib/handler"
	"main/lib/low"
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
	low.SETDATA()
	list.List(handler.Menu)
}

func Event(e model.Event) {
	function := handler.MAP[e.Key]

	if function != nil {
		function(e)
		low.LOG(e)
	} else {
		response.UNDEFINED_KEYWORD()
	}
}
