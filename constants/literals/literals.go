package literals

var FLAGS = flagsType{
	NOLOG: "-nl",
}

var Response = responseType{
	Success:          "/ᐠ-˕-マᶻ",
	UndefinedCommand: "Undefined command",
	SyntaxError:      "Syntax error",
}

var Extension = extensionType{
	Bat: ".bat",
}

var Path = pathType{
	PathLog:     "/log.asc",
	PathSetting: "/settings.asc",
}

var EventList = eventType{
	CRYPTO:   "crypto",
	DECRYPTO: "ecrypto",

	GENERATEMASTER: "gmaster",
	GENERATEKEY:    "gkey",

	CLEARLOG: "clog",
	LOG:      "log",

	DROP: "drop",
	STOP: "stop",
	HELP: "help",

	BAT:       "bat",
	RUNBAT:    "run",
	REMOVEBAT: "rbat",

	BATLIST:       "list",
	BATLISTREMOVE: "listremove",

	MENU:  "menu",
	QNEXT: "qnext",

	// ТЕХНИЧЕСКИЕ

	INWORK:       "inwork",
	RESPONSE:     "response",
	MANUADDCMD:   "MANUADDCMD",
	MENUDECRYPTO: "MENUDECRYPTO",
	MENUCRYPTO:   "MENUCRYPTO",
	RRBAT:        "%333RRBAT",
	RESET:        "reset",
}

var Titles = titleType{
	Menu:     "Меню",
	ListCmd:  "Список команд",
	EnterCmd: "Добавить команду",
	Access:   "Подтвердите действие (yea/no)",
}

var SGR = SGRtype{
	BOLD:      1,
	ITALIC:    3,
	DIM:       2,
	UNDERLINE: 4,

	RED:     31,
	GREEN:   32,
	YELLOW:  33,
	BLUE:    34,
	MAGENTA: 35,
	CYAN:    36,
	WHITE:   37,

	RESET: 0,
}
