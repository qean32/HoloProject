package handler

import (
	"atomicgo.dev/cursor"

	"main/constants"
	eventlist "main/constants/event-list"
	"main/constants/literal"
	"main/model"
	"main/terminal"
	"main/terminal/list"
)

var Menu = []model.Option{
	{Message: "Запустить команду", Event: model.Event{Key: eventlist.BATLIST}},
	{Message: "Ручной ввод", Event: model.Event{Key: eventlist.MANUAL}},
	{Message: "Удалить команду", Event: model.Event{Key: eventlist.BATLISTREMOVE}},
	{Message: "Добавить команду", Event: model.Event{Key: eventlist.MANUADDBAT}},
	{Message: "Генерация мастер ключа", Event: model.Event{Key: eventlist.INWORK}},
	{Message: "Помощь", Event: model.Event{Key: eventlist.HELP}},
	{Message: "Выход", Event: model.Event{Key: eventlist.STOP}},
}

func runMenu() {
	list.List(model.ListPayload{
		Options: Menu,
		Title:   literal.Titles.Menu,
	})
}

func help() {
	terminal.DownAndStart()
	cursor.StartOfLine()
	terminal.PrintASCIICenter(constants.HelpMessage, "")
}
