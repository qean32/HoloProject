package handler

import (
	errorlist "main/constants/error-list"
	"main/lib/handler/parse"
	"main/model"
)

func Parse(args []string) (model.Event, error) {
	if len(args) == 0 {
		return model.Event{}, errorlist.ErrEmptyCommand
	}

	command, ok := registry[args[0]]
	if !ok {
		return model.Event{}, errorlist.ErrUnknownCommand
	}
	return parse.Parse(args, command.Options)
}
