package handler

import (
	test "main/__test__"
	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
	"main/lib/death"
	"main/model"
)

var commands = []model.Command{
	{
		Name:    eventlist.BAT,
		Options: model.ParseOptions{KeyWord: true, Payload: true},
		Fn:      bat,
		Success: true,
	},
	{
		Name: eventlist.BATLIST,
	},
	{
		Name:    eventlist.RUNBAT,
		Options: model.ParseOptions{KeyWord: true},
		Fn:      runBat,
		Success: true,
	},
	{
		Name:    eventlist.REMOVEBAT,
		Options: model.ParseOptions{KeyWord: true, GetSubEvent: true, SubEvent: eventlist.PUREREMOVEBAT},
	},
	{
		Name: eventlist.MENU,
	},
	{
		Name: eventlist.HELP,
		Fn:   ignoreError(help),
	},
	{
		Name:    eventlist.CLEARLOG,
		Fn:      ignoreEvent(clearLog),
		Success: true,
	},
	{
		Name: eventlist.OPENLOG,
		Fn:   ignoreEvent(openLog),
	},
	{
		Name: eventlist.LOGGER,
		Fn:   ignoreError(toggleLogger),
	},
	{
		Name:    eventlist.DROP,
		Fn:      ignoreEvent(drop),
		Success: true,
	},
	{
		Name: eventlist.STOP,
		Fn:   ignoreError(death.Exit),
	},
	// Криптокоманды пока не реализованы
	{
		Name:    eventlist.CRYPTO,
		Options: model.ParseOptions{KeyWord: true, Payload: true, Password: true},
		Fn:      ignoreEvent(encrypt),
		Success: true,
	},
	{
		Name:    eventlist.DECRYPTO,
		Options: model.ParseOptions{KeyWord: true, Password: true},
		Fn:      ignoreEvent(decrypt),
	},
	{
		Name: eventlist.GENERATEKEY,
		Fn:   ignoreEvent(generateKey),
	},
	{
		Name: eventlist.GENERATEMASTER,
		Fn:   ignoreEvent(generateMasterKey),
	},
	{
		Name:    eventlist.TEST,
		Fn:      ignoreError(test.Test),
		Success: true,
	},
	// Служебные команды и события меню
	{
		Name: eventlist.MANUCREATEBAT,
	},
	{
		Name: eventlist.BATLISTREMOVE,
	},
	{
		Name:    eventlist.PUREREMOVEBAT,
		Fn:      removeCmd,
		Success: true,
	},
	{
		Name:    eventlist.INIT,
		Fn:      ignoreEvent(init_),
		Success: true,
	},
	{
		Name: eventlist.INWORK,
		Fn:   ignoreError(inwork),
	},
	{
		Name: eventlist.RESPONSE,
		Fn:   _response,
	},
	{
		Name: eventlist.TESTVIEW,
		Fn:   ignoreError(test.TestView),
	},
	{
		Name: eventlist.MENUCRYPTO,
		Fn:   ignoreError(inwork),
	},
	{
		Name: eventlist.RESET,
		Fn:   ignoreError(reset),
	},
}

var registry = buildRegistry(commands)

func buildRegistry(list []model.Command) map[string]model.Command {
	result := make(map[string]model.Command, len(list))
	for _, command := range list {
		if command.Name == "" {
			panic("invalid command: empty name")
		}
		if _, exists := result[command.Name]; exists {
			panic("duplicate command: " + command.Name)
		}
		result[command.Name] = command
	}
	return result
}

func ignoreError(fn func()) model.Handler {
	return func(_ model.Event) error {
		fn()
		return nil
	}
}

func ignoreEvent(fn func() error) model.Handler {
	return func(_ model.Event) error {
		return fn()
	}
}

func Dispatch(e model.Event) (model.Action, error) {
	command, ok := registry[e.Key]
	if !ok {
		return nil, errorlist.ErrUnknownCommand
	}
	return nil, command.Fn(e)
}

func Success(name string) bool {
	command, ok := registry[name]
	return ok && command.Success
}
