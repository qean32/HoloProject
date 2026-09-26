package handler

import (
	"main/constants/literals"
	"main/model"
)

var Menu = []model.Option{
	{
		Message: "Запустить команду", Event: model.Event{Key: literals.EventList.EventList},
	},
	{
		Message: "Ручной ввод", Event: model.Event{Key: literals.EventList.MANUAL},
	},
	{
		Message: "Добавить команду", Event: model.Event{Key: literals.EventList.MANUADDCMD},
	},
	{
		Message: "Шифрование строки", Event: model.Event{Key: literals.EventList.INWORK},
	},
	{
		Message: "Очистка логов", Event: model.Event{Key: literals.EventList.CLEARLOG},
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
