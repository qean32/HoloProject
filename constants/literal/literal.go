package literal

var Flags = flagType{
	NOLOG: "-nl",
}

var Extension = extensionType{
	Bat:  ".bat",
	Txt: ".txt",
}

var Path = pathType{
	PathLog: "log" + Extension.Txt,
}

var Titles = titleType{
	Menu:     "Меню",
	ListBat:  "Список команд",
	EnterBat: "Добавить команду",
	Access:   "Подтвердите действие (yea/no)",
}
