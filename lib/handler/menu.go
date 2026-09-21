package handler

import (
	"main/constants/literals"
	"main/model"
)

var Menu = []model.Option{
	{
		Message: "Запустить команду",
		Event:   model.Event{Key: literals.COMMANDS_LIST.COMMANDS_LIST},
	},
	{
		Message: "Ручной ввод",
		Event:   model.Event{Key: literals.COMMANDS_LIST.MANUAL},
	},
	{
		Message: "Добавить команду",
		Event:   model.Event{Key: literals.COMMANDS_LIST.MANU_ADD_COMMAND},
	},
	{
		Message: "Шифрование строки",
		Event:   model.Event{Key: literals.COMMANDS_LIST.INWORK},
	},
	{
		Message: "Очистка логов",
		Event:   model.Event{Key: literals.COMMANDS_LIST.CLEARLOG},
	},
	{
		Message: "Генерация мастер ключа",
		Event:   model.Event{Key: literals.COMMANDS_LIST.INWORK},
	},
	{
		Message: "Помощь",
		Event:   model.Event{Key: literals.COMMANDS_LIST.HELP},
	},
	{
		Message: "Выход",
		Event:   model.Event{Key: literals.COMMANDS_LIST.STOP},
	},
}
