package handler

import (
	"main/callstack/manual"
	eventlist "main/constants/event-list"
	"main/lib/death"
	"main/model"
	qnext "main/terminal/qNext"
)

var registry = map[string]model.Handler{
	// клиентские
	eventlist.INIT: init_,

	eventlist.BAT:        bat,
	eventlist.MANUADDBAT: questionnaireAddBat,

	eventlist.RUNBAT:    runBat,
	eventlist.REMOVEBAT: questionnaireAccess,

	eventlist.BATLIST:       runBatList,
	eventlist.BATLISTREMOVE: removeBatList,

	eventlist.HELP:   ignoreError(help),
	eventlist.MENU:   ignoreError(runMenu),
	eventlist.RESET:  reset,
	eventlist.DROP:   drop,
	eventlist.STOP:   ignoreError(death.Exit),
	eventlist.MANUAL: ignoreError(manual.Manual),

	eventlist.CLEARLOG: clearLog,
	eventlist.OPENLOG:  openLog,

	eventlist.CRYPTO:   encrypt,
	eventlist.DECRYPTO: decrypt,

	eventlist.GENERATEKEY:    generateKey,
	eventlist.GENERATEMASTER: generateMasterKey,
	// технические
	eventlist.RRBAT:      removeCmd,
	eventlist.RESPONSE:   _response,
	eventlist.MENUCRYPTO: inwork,
	eventlist.QNEXT:      ignoreError(qnext.QNext),
	eventlist.MANUADDBAT: questionnaireAddBat,
}

func ignoreError(fn func()) model.Handler {
	return func(_ model.Event) error {
		fn()
		return nil
	}
}

func Dispatch(e model.Event) error {
	fn, ok := registry[e.Key]
	if !ok {
		return ErrUnknownCommand
	}
	return fn(e)
}
