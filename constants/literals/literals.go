package literals

var FLAGS = flagsType{
	NOLOG: "-nl",
}

var COMMANDLIST = commandsType{
	CRYPTO:   "crypto",
	DECRYPTO: "ecrypto",

	GENERATEMASTER: "gmaster",
	GENERATEKEY:    "gkey",

	CLEARLOG: "clog",
	LOGS:     "logs",

	DROP:   "drop",
	STOP:   "stop",
	HELP:   "help",
	INWORK: "inwork",

	DECLARE:            "declare",
	COMMANDLIST:        "clist",
	RUNCOMMAND:         "run",
	RUNMULTIPLECOMMAND: "runm",
	REMOVECOMMAND:      "rmc",

	MENU: "menu",

	MANUADDCOMMAND: "MANUADDCOMMAND",
	MENUDECRYPTO:   "MENUDECRYPTO",
	MENUCRYPTO:     "MENUCRYPTO",
}
