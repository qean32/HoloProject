package literal

var Flags = flagType{
	NOLOG: "-nl",
}

var Extension = extensionType{
	Bat: ".bat",
}

var Path = pathType{
	PathLog:     "log.asc",
	PathSetting: "settings.asc",
}

var Titles = titleType{
	Menu:     "Меню",
	ListBat:  "Список команд",
	EnterBat: "Добавить команду",
	Access:   "Подтвердите действие (yea/no)",
}
