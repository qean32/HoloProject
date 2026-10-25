package registry

import (
	test "main/__test__"
	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
	"main/lib/death"
	"main/lib/handler"
	"main/lib/handler/menu"
	"main/model"
)

var commands = []model.Command{
	{
		Name:    eventlist.BAT,
		Options: model.ParseOptions{KeyWord: true, Payload: true},
		Fn:      handler.CreateBat,
		Success: true,
	},
	{
		Name:    eventlist.RUNBAT,
		Options: model.ParseOptions{KeyWord: true},
		Fn:      handler.RunBat,
		Success: true,
	},
	{
		Name:       eventlist.REMOVEBAT,
		Options:    model.ParseOptions{SubEvent: eventlist.PUREREMOVEBAT, GetSubEvent: true},
		NeedAccess: true,
		Fn:         handler.RemoveCmd,
	},
	{
		Name:    eventlist.TEST,
		Fn:      ignoreError(test.Test),
		Success: true,
	},
	{Name: eventlist.DROP, Fn: ignoreEvent(handler.Drop), Success: true, NeedAccess: true},
	{Name: eventlist.CLEARLOG, Fn: ignoreEvent(handler.ClearLog), Success: true},
	{Name: eventlist.OPENLOG, Fn: ignoreEvent(handler.ClearLog)},
	{Name: eventlist.LOGGER, Fn: ignoreError(handler.ToggleLogger)},
	{Name: eventlist.STOP, Fn: ignoreError(death.Exit)},
	{Name: eventlist.MENU, Fn: ignoreError(menu.MainMenu)},
	{Name: eventlist.HELP, Fn: ignoreError(handler.Help)},
	{Name: eventlist.BATLIST, Fn: ignoreError(func() { menu.BatList(eventlist.RUNBAT) })},
	{Name: eventlist.BATLISTREMOVE, Fn: ignoreError(func() { menu.BatList(eventlist.REMOVEBAT) })},
	// Служебные команды и события меню
	{
		Name:       eventlist.PUREREMOVEBAT,
		Fn:         handler.RemoveCmd,
		Success:    true,
	},
	{
		Name:       eventlist.INIT,
		Fn:         ignoreEvent(handler.InitSystem),
		Success:    true,
		NeedAccess: true,
	},
	{Name: eventlist.INWORK, Fn: ignoreError(handler.Inwork)},
	{Name: eventlist.RESPONSE, Fn: handler.Response},
	{Name: eventlist.TESTVIEW, Fn: ignoreError(test.TestView)},
	{Name: eventlist.RESET, Fn: ignoreError(handler.Reset)},
}

var registry = buildRegistry(commands)

func Dispatch(e model.Event) (model.Action, error) {
	command, ok := registry[e.Key]

	if command.Options.GetSubEvent {
		e.SubEvent = &model.Event{Key: command.Options.SubEvent, KeyWord: e.KeyWord}
	}

	if !ok {
		return nil, errorlist.ErrUnknownCommand
	}
	if command.NeedAccess {
		if err := handler.QAccess(e, ""); err != nil {
			return nil, err
		}
	}
	return nil, command.Fn(e)
}
