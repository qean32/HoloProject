package menu

import (
	"main/lib"
	"main/model"
)

var Menu = []model.Option{
	{
		Message: "Запустить команду",
		Command: runCommand,
	},
	{
		Message: "Ручной ввод",
		Command: manual,
	},
	{
		Message: "Добавить команду",
		Command: addCommand,
	},
	{
		Message: "Очистка логов",
		Command: lib.ClearLog,
	},
	{
		Message: "Генерация мастер ключа",
		Command: lib.Help,
	},
	{
		Message: "Помощь",
		Command: lib.Help,
	},
	{
		Message: "Выход",
		Command: lib.StopProcess,
	},
}
