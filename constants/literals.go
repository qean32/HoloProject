package constants

type commandsType struct {
	CRIPTO  string
	ECRIPTO string

	GMASTER string
	GKEY    string

	CLOG string
	DROP string
	STOP string
	HELP string

	NOTE  string
	DNOTE string
	NOTES string

	DECLARE  string
	COMMANDS string
	RUN      string
	RUNM     string
	RMC      string
}

type flagsType struct {
	NOLOG string
}

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
