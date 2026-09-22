package manual

import (
	"main/callstack"
	"main/constants"
	"main/lib/parse"
	"main/terminal"
	"main/terminal/field"
	"strings"
)

func Manual() {
	command := field.Field()

	if len(command) > 1 {
		trimString := strings.TrimSpace(command)
		event, _error := parse.ParseEvent(trimString, strings.Split(trimString, " ")[0])

		if !_error {
			callstack.PushCallStack(event)
		} else {
			terminal.Outputln(constants.SYNTAX_ERROR)
		}
	}
}
