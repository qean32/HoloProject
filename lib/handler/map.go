package handler

import (
	"main/callstack/manual"
	"main/constants/literals"
	"main/lib/death"
	"main/model"
)

var MAP = map[string]model.EventFunction{
	literals.EventList.CRYPTO:   encrypt,
	literals.EventList.DECRYPTO: decrypt,

	literals.EventList.GENERATEKEY:    generateKey,
	literals.EventList.GENERATEMASTER: generateMasterKey,

	literals.EventList.DECLARE:   declare,
	literals.EventList.RUNCMD:    runCmd,
	literals.EventList.REMOVECMD: removeCmd,

	literals.EventList.CLEARLOG: clearLog,
	literals.EventList.LOGS:     ignoreEvent(openLog),

	literals.EventList.RESPONSE: _response,

	literals.EventList.MENU: runMenu,
	literals.EventList.DROP: drop,
	literals.EventList.STOP: ignoreEvent(death.Exit),
	literals.EventList.HELP: help,

	literals.EventList.EventList:  ignoreEvent(cmdList),
	literals.EventList.MANUAL:     ignoreEvent(manual.Manual),
	literals.EventList.MANUADDCMD: questionnaireAddCommand,
	literals.EventList.MENUCRYPTO: inwork,
}

func ignoreEvent(fn func()) model.EventFunction {
	return func(model.Event) { fn() }
}
