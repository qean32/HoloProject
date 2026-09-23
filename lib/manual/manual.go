package manual

import (
	"strings"

	"main/callstack"
	"main/constants/response"
	"main/lib/parse"
	"main/terminal/field"
)

func Manual() {
	command := field.Field()
	if len(command) <= 1 {
		return
	}

	trimmed := strings.TrimSpace(command)
	if trimmed == "" {
		return
	}

	event, hasError := parse.ParseEvent(trimmed, strings.Split(trimmed, " ")[0])
	if hasError {
		response.SyntaxError()
		return
	}
	callstack.PushCallStack(event)
}
