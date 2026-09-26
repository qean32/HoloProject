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
	LOGS:     "logs",

	DROP:   "drop",
	STOP:   "stop",
	HELP:   "help",
	INWORK: "inwork",

	DECLARE:   "declare",
	EventList: "clist",
	RUNCMD:    "run",
	REMOVECMD: "rmc",

	RESPONSE: "response",

	MENU: "menu",

	MANUADDCMD:   "MANUADDCMD",
	MENUDECRYPTO: "MENUDECRYPTO",
	MENUCRYPTO:   "MENUCRYPTO",
}

var Titles = titleType{
	Menu:     "Меню",
	ListCmd:  "Список команд",
	EnterCmd: "Добавить команду",
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
