package handler

import (
	test "main/__test__"
	"main/callstack/manual"
	"main/lib/death"
	"main/model"

	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
)

var registry = map[string]model.Handler{
	// клиентские
	eventlist.INIT: ignoreEvent(init_),

	eventlist.BAT:           bat,
	eventlist.MANUCREATEBAT: ignoreEvent(questionnaireCreateBat),

	eventlist.RUNBAT:    runBat,
	eventlist.REMOVEBAT: questionnaireAccess,

	eventlist.BATLIST:       ignoreEvent(runBatList),
	eventlist.BATLISTREMOVE: ignoreEvent(removeBatList),

	eventlist.HELP:   ignoreError(help),
	eventlist.MENU:   ignoreError(runMenu),
	eventlist.RESET:  ignoreError(reset),
	eventlist.DROP:   ignoreEvent(drop),
	eventlist.STOP:   ignoreError(death.Exit),
	eventlist.MANUAL: ignoreError(manual.Manual),

	eventlist.LOGGER: ignoreError(toggleLogger),

	eventlist.CLEARLOG: ignoreEvent(clearLog),
	eventlist.OPENLOG:  ignoreEvent(openLog),

	eventlist.CRYPTO:   ignoreEvent(encrypt),
	eventlist.DECRYPTO: ignoreEvent(decrypt),

	eventlist.GENERATEKEY:    ignoreEvent(generateKey),
	eventlist.GENERATEMASTER: ignoreEvent(generateMasterKey),
	// технические
	eventlist.PUREREMOVEBAT: removeCmd,
	eventlist.INWORK:        ignoreError(inwork),
	eventlist.TEST:          ignoreError(test.Test),
	eventlist.TESTVIEW:      ignoreError(test.TestView),
	eventlist.RESPONSE:      _response,
	eventlist.MENUCRYPTO:    ignoreError(inwork),
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

func Dispatch(e model.Event) error {
	fn, ok := registry[e.Key]
	if !ok {
		return errorlist.ErrUnknownCommand
	}
	return fn(e)
}
