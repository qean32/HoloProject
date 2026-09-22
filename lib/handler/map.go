package handler

import (
	"main/constants/literals"
	"main/lib/manual"
	"main/model"
)

var MAP = map[string]model.EventFunction{
	literals.COMMANDS_LIST.CRIPTO:  Encrypt,
	literals.COMMANDS_LIST.ECRIPTO: Decrypt,

	literals.COMMANDS_LIST.GENERATEKEY:    GenerateKey,
	literals.COMMANDS_LIST.GENERATEMASTER: GenerateMasterKey,

	literals.COMMANDS_LIST.DECLARE:               Declare,
	literals.COMMANDS_LIST.RUN_COMMAND:           RunCommand,
	literals.COMMANDS_LIST.REMOVE_COMMAND:        RemoveCommand,
	literals.COMMANDS_LIST.RUN_MULTIPLE_COMMANDS: RunMultipleCommands,

	literals.COMMANDS_LIST.CLEARLOG: ClearLog,
	literals.COMMANDS_LIST.MENU:     RunMenu,
	literals.COMMANDS_LIST.DROP:     Drop,
	literals.COMMANDS_LIST.STOP:     Stop,
	literals.COMMANDS_LIST.HELP:     Help,

	literals.COMMANDS_LIST.COMMANDS_LIST:    func(e model.Event) { Menu_runCommandsList() },
	literals.COMMANDS_LIST.MANUAL:           func(e model.Event) { manual.Manual() },
	literals.COMMANDS_LIST.MANU_ADD_COMMAND: Menu_runQuestionnaireAddCommand,
	literals.COMMANDS_LIST.MENU_HASH:        Inwork,
}
