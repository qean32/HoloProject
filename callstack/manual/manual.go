package manual

import (
	"strings"

	"main/callstack"
	"main/callstack/event"
	"main/constants"
	"main/lib/handler/parse"
	"main/model"
	"main/terminal/field"
)

func Manual() {
	command := field.Field(model.FieldPayload{Prefix: constants.FIELDPREFIX})
	trimmed := strings.TrimSpace(command)

	if len(trimmed) <= 1 {
		event.SyntaxError()
		return
	}

	_event, hasError := parse.ParseEvent(trimmed, strings.Split(trimmed, " ")[0])
	if hasError {
		event.SyntaxError()
		return
	}
	callstack.PushCallStack(_event)
}
