package lib

import (
	"bufio"
	"fmt"
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
		fmt.Println(constants.UNDEFINED_COMMAND)
		print("> ")
	}
}

var READER = bufio.NewReader(os.Stdin)

func ENTER_COMMAND(callstack model.Channel) {
	print("> ")
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
	fmt.Println(constants.PROJECT_INIT)
	deep.SET_DATA()
	ENTER_COMMAND(callstack)
}

func PARSE_EVENT(command string, key string) (e model.Event, _error bool) {
	fn := KEY_PARSE[key]

	if fn == nil {
		e, _error = SHORT_EVENT(strings.Split(command, " "))
		return
	}
	e, _error = fn(strings.Split(command, " "))
	return e, _error
}
