package manual

import (
	"strings"

	"main/callstack"
	"main/constants"
	"main/constants/response"
	"main/lib/parse"
	"main/model"
	"main/terminal/field"
)

func Manual() {
	command := field.Field(model.FieldPayload{Prefix: constants.FIELDPREFIX})
	if len(command) <= 1 {
		response.SyntaxError()
		return
	}

	trimmed := strings.TrimSpace(command)
	if trimmed == "" {
		response.SyntaxError()
		return
	}

	event, hasError := parse.ParseEvent(trimmed, strings.Split(trimmed, " ")[0])
	if hasError {
		response.SyntaxError()
		return
	}
	callstack.PushCallStack(event)
}
