package lib

import (
	"bufio"
	"main/constants"
	"main/deep"
	"main/model"
	"os"
	"strings"
)

func ITERATION_CYCLE(e model.Event) {
	fn := KEY_FUNCTION[e.Key]

	if fn != nil {
		fn(e)
		deep.LOG(e)
	} else {
		deep.Console(constants.UNDEFINED_COMMAND)
	}
}

var READER = bufio.NewReader(os.Stdin)

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

func INIT(callstack model.Channel) {
	deep.Console(constants.PROJECT_INIT)
	deep.SET_DATA()
	ENTER_COMMAND(callstack)
}

func matrixToArrayString(matrix [][]string) []string {
	var tmpArr []string

	for i := 0; i < len(matrix); i++ {
		tmpArr = append(tmpArr, strings.Join(matrix[i], " "))
	}

	return tmpArr
}
