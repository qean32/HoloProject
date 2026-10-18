package manual

import (
	"strings"

	"main/callstack"
	"main/callstack/event"
	"main/constants"
	errorlist "main/constants/error-list"
	"main/lib/handler/parse"
	"main/model"
	"main/terminal/field"
)

func Manual() {
	command := field.Field(model.FieldPayload{Prefix: constants.FIELDPREFIX})
	args := strings.Fields(command)
	if len(args) == 0 {
		event.Response(errorlist.ErrEmptyCommand)
		return
	}

	e, err := parse.Dispatch(args[0], args)
	if err != nil {
		event.Response(err)
		return
	}
	callstack.Push(e)
}
