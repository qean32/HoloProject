package lib

import (
	"main/callstack"
	"main/callstack/event"
	"main/constants"
	"main/constants/literal"
	"main/lib/death"
	"main/lib/handler"
	"main/model"
	"main/terminal"
)

func Init() error {
	if err := InitData(); err != nil {
		return err
	}

	terminal.RenderBanner()
	return nil
}

func InitData() error {
	if err := constants.InitRoot(); err != nil {
		return err
	}
	return death.EnsureFile(literal.Path.PathLog)
}

func HandleEvent(e model.Event) error {
	action, err := handler.Dispatch(e)
	if err != nil {
		return fail(e, err)
	}

	if action != nil {
		events, err := action.Run()
		if err != nil {
			return fail(e, err)
		}
		for _, next := range events {
			callstack.Push(next)
		}
	}

	if handler.Success(e.Key) {
		event.Success()
	}

	return logg(e)
}

// logg ошибки
func fail(e model.Event, err error) error {
	event.Response(err)
	if logErr := death.Logger(model.StackError{Event: e, Err: err}); logErr != nil {
		event.Response(logErr)
	}
	return err
}

func logg(e model.Event) error {
	if err := death.Logger(model.StackError{Event: e}); err != nil {
		event.Response(err)
		return err
	}
	return nil
}
