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

func INIT(callstack model.Channel) {
	deep.Console(constants.PROJECT_INIT)
	deep.SET_DATA()
	ENTER_COMMAND(callstack)
}

func ITERATION_CYCLE(e model.Event) {
	fn := KEY_FUNCTION[e.Key]

	if fn != nil {
		fn(e)
		deep.LOG(e)
	} else {
		deep.Console(constants.UNDEFINED_COMMAND)
	}
}

func ENTER_COMMAND(callstack model.Channel) {
	command, _ := READER.ReadString('\n')

	if len(command) > 1 {
		trimString := strings.TrimSpace(command)
		key := strings.Split(trimString, " ")[0]
		e, _error := PARSE_EVENT(trimString, key)

		if !_error {
			deep.CALLSTACK = append(deep.CALLSTACK, e)
			callstack <- e
		}
	}
	ENTER_COMMAND(callstack)
}
