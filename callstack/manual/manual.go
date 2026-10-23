package manual

import (
	"errors"
	"io"
	"strings"

	"main/callstack"
	"main/callstack/event"
	"main/constants"
	"main/lib/death"
	"main/lib/handler"
	"main/model"
	"main/terminal/field"
)

func Manual() {
	command, err := field.Field(model.FieldPayload{Prefix: constants.FIELDPREFIX})
	if err != nil {
		if errors.Is(err, io.EOF) {
			death.Exit()
		}
		event.Response(err)
		return
	}

	args := strings.Fields(command)
	e, err := handler.Parse(args)
	if err != nil {
		event.Response(err)
		return
	}
	callstack.Push(e)
}
