package lib

import (
	"bufio"
	"os"

	"main/constants"
	"main/constants/literals"
	"main/constants/response"
	"main/lib/handler"
	"main/lib/low"
	"main/model"
	"main/terminal"
	"main/terminal/list"
)

var READER = bufio.NewReader(os.Stdin)

func INIT() {
	terminal.PrintASCIICenter(constants.BinaryPROJECT_INIT, " ")
	terminal.PrintASCIICenter(constants.PROJECT_INIT, " ")
	constants.INIT_ROOT()
	low.SETDATA()
	list.List(handler.Menu, literals.Titles.Menu)
}

func Event(e model.Event) {
	defer response.Success()
	function := handler.MAP[e.Key]

	if function == nil {
		response.UndefinedKeyword()
		return
	}
	function(e)
	low.LOG(e)
}
