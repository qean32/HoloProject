package lib

import (
	"bufio"
	"main/constants"
	"main/lib/low"
	"main/model"
	"main/terminal"
	"main/terminal/field"
	"os"
	"strings"
)

var READER = bufio.NewReader(os.Stdin)

func INIT() {
	terminal.OutputASCII_CENTER(constants.BinaryPROJECT_INIT, " ")
	terminal.OutputASCII_CENTER(constants.PROJECT_INIT, " ")
	constants.INIT_ROOT()
	low.SETDATA()

	MANUAL()
}

func ITERATION_CYCLE(e model.Event) {
	function := MAP_HANDLER[e.Key]

	if function != nil {
		function(e)
		low.LOG(e)
	} else {
		terminal.Outputln(constants.UNDEFINED_COMMAND)
	}
}

func Help() {
	terminal.OutputASCII_CENTER(constants.HelpMessage, "")
}

func RunCommand(command string) {
	low.RUN_CMD(command)
}

func MANUAL() {
	command := field.Field()

	if len(command) > 1 {
		trimString := strings.TrimSpace(command)
		event, _error := PARSE_EVENT(trimString, strings.Split(trimString, " ")[0])

		if !_error {
			low.CALLSTACK = append(low.CALLSTACK, event)
			ITERATION_CYCLE(event)
		} else {
			terminal.Outputln(constants.SYNTAX_ERROR)
		}
	}
}

func InWork() {
	terminal.Outputln("В разработке")
}
