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

var READER = bufio.NewReader(os.Stdin)

func ENTER_COMMAND() {
	print("> ")
	command, _ := READER.ReadString('\n')

	if len(command) > 1 {
		trimString := strings.TrimSpace(command)
		key := strings.Split(trimString, " ")[0]
		event, _error := PARSE_EVENT(trimString, key)

		if !_error {
			fn := KEY_FUNCTION[key]

			if fn != nil {
				fn(event)
				deep.LOG(event)
			} else {
				fmt.Println(constants.UNDEFINED_COMMAND)
			}
		}
	}
	ENTER_COMMAND()
}

func INIT() {
	fmt.Println(constants.PROJECT_INIT)
	deep.SET_DATA()
	ENTER_COMMAND()
}

func PUSH_CYCLE() {
}

func PARSE_EVENT(command string, key string) (event model.Event, _error bool) {
	fn := KEY_PARSE[key]

	if fn == nil {
		event, _error = SHORT_EVENT(strings.Split(command, " "))
		return
	}
	event, _error = fn(strings.Split(command, " "))
	return event, _error
}
