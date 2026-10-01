package handler

import (
	"main/callstack/manual"
	"main/constants/literals"
	"main/lib/death"
	"main/model"
)

var MAP = map[string]model.EventFunction{
	literals.EventList.CRYPTO:   inwork,
	literals.EventList.DECRYPTO: inwork,

	literals.EventList.GENERATEKEY:    inwork,
	literals.EventList.GENERATEMASTER: inwork,

	literals.EventList.BAT:       bat,
	literals.EventList.RUNBAT:    runCmd,
	literals.EventList.REMOVEBAT: questionnaireAccess,
	literals.EventList.RRBAT:     removeCmd,

	literals.EventList.CLEARLOG: clearLog,
	literals.EventList.LOG:      ignoreEvent(openLog),

	literals.EventList.RESPONSE: _response,

	literals.EventList.MENU: runMenu,
	literals.EventList.DROP: drop,
	literals.EventList.STOP: ignoreEvent(death.Exit),
	literals.EventList.HELP: help,

	literals.EventList.BATLIST: ignoreEvent(
		func() {
			batList(literals.EventList.RUNBAT)
		},
	),
	literals.EventList.BATLISTREMOVE: ignoreEvent(removeBatList),
	literals.EventList.MANUAL:        ignoreEvent(manual.Manual),
	literals.EventList.MANUADDCMD:    questionnaireAddCommand,
	literals.EventList.MENUCRYPTO:    inwork,
	literals.EventList.RESET:         ignoreEvent(reset),
	literals.EventList.QNEXT:         ignoreEvent(_qnext),
}

func ignoreEvent(fn func()) model.EventFunction {
	return func(model.Event) { fn() }
}
