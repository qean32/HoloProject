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
	stlit := strings.Split(strings.TrimSpace(command), " ")

	e, err := parse.Dispatch(stlit[0], stlit)
	if err != nil {
		event.Response(err)
		return
	}
	callstack.Push(e)
}
