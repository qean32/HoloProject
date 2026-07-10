package literals

var FLAGS = flagsType{
	NOLOG: "-nl",
}

/*
 CRIPTO добавить сохранение даты
 CRIPTO добавить шифрование
 NOTE сдеать записную строку сохранение даты шифрование пароль
*/

var COMMANDS = commandsType{
	CRIPTO:  "cripto",
	ECRIPTO: "ecripto",

	GMASTER: "gmaster",
	GKEY:    "gkey",

	CLOG: "clog",
	DROP: "drop",
	STOP: "stop",
	HELP: "help",

	NOTE:  "note",
	DNOTE: "dnote",
	NOTES: "notes",

	DECLARE:  "declare",
	COMMANDS: "comm",
	RUN:      "run",
	RUNM:     "runm",
	RMC:      "rmc",
}
