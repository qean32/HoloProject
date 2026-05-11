package lib

import (
	"bufio"
	"main/constants"
	"main/deep"
	"main/model"
	"os"
	"strings"
)

var READER = bufio.NewReader(os.Stdin)

func INIT(callstack_channel model.CallStackChannel) {
	deep.CONSOLE_ASCII_CENTER(constants.PROJECT_INIT, "")
	constants.INIT_ROOT()
	deep.DATA()
	ENTER_COMMAND(callstack_channel)
}

func ITERATION_CYCLE(e model.Event) {
	function := MAP_HANDLER[e.Key]

	if function != nil {
		function(e)
		deep.CONSOLE("")
		deep.LOG(e)
	} else {
		deep.CONSOLE_RESPONSE(constants.UNDEFINED_COMMAND, true)
	}
}

func ENTER_COMMAND(callstack_channel model.CallStackChannel) {
	command, _ := READER.ReadString('\n')

	if len(command) > 1 {
		trimString := strings.TrimSpace(command)
		key := strings.Split(trimString, " ")[0]
		e, _error := PARSE_EVENT(trimString, key)

		if !_error {
			deep.CALLSTACK = append(deep.CALLSTACK, e)
			// callstack_channel <- e
			ITERATION_CYCLE(e)
		} else {
			deep.CONSOLE_RESPONSE(constants.SYNTAX_ERROR, true)
		}
	}
	ENTER_COMMAND(callstack_channel)
}

func get_it_out_of_the_basket(callstack_channel model.CallStackChannel) model.Event {
	it := deep.CALLSTACK[:len(deep.CALLSTACK)-1][0]
	deep.CALLSTACK = deep.CALLSTACK[:len(deep.CALLSTACK)-1]
	callstack_channel <- it

	return it
}
