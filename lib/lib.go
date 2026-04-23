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

func INIT(callstack_channel model.CallStackChannel) {
	deep.CONSOLE(constants.PROJECT_INIT)
	deep.DATA()
	ENTER_COMMAND(callstack_channel)
}

func ITERATION_CYCLE(e model.Event) {
	fn := KEY_FUNCTION[e.Key]

	if fn != nil {
		fn(e)
		deep.CONSOLE("")
		deep.LOG(e)
	} else {
		deep.CONSOLE(constants.UNDEFINED_COMMAND)
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
			callstack_channel <- e
		}
	}
	ENTER_COMMAND(callstack_channel)
}

func get_it_out_of_the_basket(callstack_channel model.CallStackChannel) model.Event {
	it := deep.CALLSTACK[:len(deep.CALLSTACK)-1][0]
	fmt.Println(it)
	deep.CALLSTACK = deep.CALLSTACK[:len(deep.CALLSTACK)-1]
	callstack_channel <- it

	return it
}
