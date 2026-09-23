package handler

import (
	"main/constants/literals"
	"main/lib/low"
	"main/lib/manual"
	"main/model"
)

var MAP = map[string]model.EventFunction{
	literals.COMMANDLIST.CRYPTO:   encrypt,
	literals.COMMANDLIST.DECRYPTO: decrypt,

	literals.COMMANDLIST.GENERATEKEY:    generateKey,
	literals.COMMANDLIST.GENERATEMASTER: generateMasterKey,

	literals.COMMANDLIST.DECLARE:            declare,
	literals.COMMANDLIST.RUNCOMMAND:         runCommand,
	literals.COMMANDLIST.REMOVECOMMAND:      removeCommand,
	literals.COMMANDLIST.RUNMULTIPLECOMMAND: runMultipleCommand,

	literals.COMMANDLIST.CLEARLOG: clearLog,
	literals.COMMANDLIST.LOGS:     ignoreEvent(openLog),

	literals.COMMANDLIST.MENU: runMenu,
	literals.COMMANDLIST.DROP: drop,
	literals.COMMANDLIST.STOP: ignoreEvent(low.Exit),
	literals.COMMANDLIST.HELP: help,

	literals.COMMANDLIST.COMMANDLIST:    ignoreEvent(menuCommandList),
	literals.COMMANDLIST.MANUAL:         ignoreEvent(manual.Manual),
	literals.COMMANDLIST.MANUADDCOMMAND: menuQuestionnaireAddCommand,
	literals.COMMANDLIST.MENUCRYPTO:     inwork,
}

func ignoreEvent(fn func()) model.EventFunction {
	return func(model.Event) { fn() }
}
