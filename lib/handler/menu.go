package handler

import (
	"main/constants/literals"
	"main/model"
)

var Menu = []model.Option{
	{
		Message: "Запустить команду", Event: model.Event{Key: literals.EventList.EVENTLIST},
	},
	{
		Message: "Ручной ввод", Event: model.Event{Key: literals.EventList.MANUAL},
	},
	{
		Message: "Удалить команду", Event: model.Event{Key: literals.EventList.EVENTLISTREMOVE},
	},
	{
		Message: "Добавить команду", Event: model.Event{Key: literals.EventList.MANUADDCMD},
	},
	{
		Message: "Генерация мастер ключа", Event: model.Event{Key: literals.EventList.INWORK},
	},
	{
		Message: "Помощь", Event: model.Event{Key: literals.EventList.HELP},
	},
	{
		Message: "Выход", Event: model.Event{Key: literals.EventList.STOP},
	},
}
