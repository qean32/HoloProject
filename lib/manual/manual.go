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
	trimmed := strings.TrimSpace(command)

	if len(trimmed) <= 1 {
		callstack.PushCallStack(response.SyntaxError())
		return
	}

	_event, hasError := parse.ParseEvent(trimmed, strings.Split(trimmed, " ")[0])
	if hasError {
		callstack.PushCallStack(response.SyntaxError())
		return
	}
	callstack.PushCallStack(_event)
}
