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
	literals.COMMANDS_LIST.COMMANDS_LIST:         ListCommands,
	literals.COMMANDS_LIST.REMOVE_COMMAND:        RemoveCommand,
	literals.COMMANDS_LIST.RUN_MULTIPLE_COMMANDS: RunMultipleCommands,

	literals.COMMANDS_LIST.NOTES:      Notes,
	literals.COMMANDS_LIST.NOTE:       Note,
	literals.COMMANDS_LIST.DELETENOTE: DeleteNote,

	literals.COMMANDS_LIST.CLEARLOG: ClearLog,
	literals.COMMANDS_LIST.MENU:     RunMenu,
	literals.COMMANDS_LIST.INWORK:   Inwork,
	literals.COMMANDS_LIST.DROP:     Drop,
	literals.COMMANDS_LIST.STOP:     Stop,
	literals.COMMANDS_LIST.HELP:     Help,

	literals.COMMANDS_LIST.COMMANDS_LIST:    ListCommands,
	literals.COMMANDS_LIST.MANUAL:           manual.Manual,
	literals.COMMANDS_LIST.MANU_ADD_COMMAND: Menu_runQuestionnaireAddCommand,
	literals.COMMANDS_LIST.MENU_HASH_STRING: Inwork,
}
