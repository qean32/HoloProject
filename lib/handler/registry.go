package handler

import (
	test "main/__test__"
	"main/callstack/manual"
	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
	"main/lib/death"
	"main/model"
	qnext "main/terminal/qNext"
)

var registry = map[string]model.Handler{
	// клиентские
	eventlist.INIT: ignoreError(init_),

	eventlist.BAT:        bat,
	eventlist.MANUADDBAT: ignoreEvent(questionnaireAddBat),

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

	eventlist.CLEARLOG: ignoreEvent(clearLog),
	eventlist.OPENLOG:  ignoreEvent(openLog),

	eventlist.CRYPTO:   ignoreEvent(encrypt),
	eventlist.DECRYPTO: ignoreEvent(decrypt),

	eventlist.GENERATEKEY:    ignoreEvent(generateKey),
	eventlist.GENERATEMASTER: ignoreEvent(generateMasterKey),
	// технические
	eventlist.PURE_REMOVE_BAT: removeCmd,
	eventlist.TEST:            ignoreEvent(test.Test),
	eventlist.RESPONSE:        _response,
	eventlist.MENUCRYPTO:      ignoreError(inwork),
	eventlist.QNEXT:           ignoreError(qnext.QNext),
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
