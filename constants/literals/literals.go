package literals

var FLAGS = flagsType{
	NOLOG: "-nl",
}

/*
 CRIPTO добавить сохранение даты
 CRIPTO добавить шифрование
 NOTE сделать записную строку сохранение даты шифрование пароль
*/

var COMMANDS_LIST = commandsType{
	CRIPTO:  "cripto",
	ECRIPTO: "ecripto",

	GENERATEMASTER: "gmaster",
	GENERATEKEY:    "gkey",

	CLEARLOG: "clog",
	DROP:     "drop",
	STOP:     "stop",
	HELP:     "help",
	INWORK:   "inwork",

	DECLARE:               "declare",
	COMMANDS_LIST:         "clist",
	RUN_COMMAND:           "run",
	RUN_MULTIPLE_COMMANDS: "runm",
	REMOVE_COMMAND:        "rmc",

	MENU: "menu",

	MANU_ADD_COMMAND: "MANU_ADD_COMMAND",
	MENU_HASH:        "MENU_HASH",
}
