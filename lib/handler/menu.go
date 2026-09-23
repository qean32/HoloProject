package handler

import (
	"main/constants/literals"
	"main/model"
)

var Menu = []model.Option{
	{
		Message: "Запустить команду", Event: model.Event{Key: literals.COMMANDLIST.COMMANDLIST},
	},
	{
		Message: "Ручной ввод", Event: model.Event{Key: literals.COMMANDLIST.MANUAL},
	},
	{
		Message: "Добавить команду", Event: model.Event{Key: literals.COMMANDLIST.MANUADDCOMMAND},
	},
	{
		Message: "Шифрование строки", Event: model.Event{Key: literals.COMMANDLIST.INWORK},
	},
	{
		Message: "Очистка логов", Event: model.Event{Key: literals.COMMANDLIST.CLEARLOG},
	},
	{
		Message: "Генерация мастер ключа", Event: model.Event{Key: literals.COMMANDLIST.INWORK},
	},
	{
		Message: "Помощь", Event: model.Event{Key: literals.COMMANDLIST.HELP},
	},
	{
		Message: "Выход", Event: model.Event{Key: literals.COMMANDLIST.STOP},
	},
}
