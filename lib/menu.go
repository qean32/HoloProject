package lib

import (
	"main/constants/literals"
	"main/lib/array"
	"main/lib/low"
	"main/model"
	"main/terminal/list"
	"main/terminal/questionnaire"
)

var Menu = []model.Option{
	{
		Message: "Запустить команду",
		Command: commands,
	},
	{
		Message: "Ручной ввод",
		Command: MANUAL,
	},
	{
		Message: "Добавить команду",
		Command: addCommand,
	},
	{
		Message: "Шифрование строки",
		Command: addCommand,
	},
	{
		Message: "Очистка логов",
		Command: low.ClearLog,
	},
	{
		Message: "Генерация мастер ключа",
		Command: InWork,
	},
	{
		Message: "Помощь",
		Command: Help,
	},
	{
		Message: "Выход",
		Command: low.StopProcess,
	},
}

func commands() {
	list.List(
		array.Map(low.TMP_COMMANDS, func(value []string) model.Option {
			return model.Option{
				Message: value[0],
				Command: func() {
					RunCommand(value[1])
				},
			}
		}),
	)
}

func addCommand() {
	q := questionnaire.Questionnaire([]model.Question{
		{
			Message: "Ключ",
			Key:     "KeyWord",
			Callback: func(res string) bool {
				return true
			},
		},
		{
			Message: "Команда",
			Key:     "Payload",
			Callback: func(res string) bool {
				return true
			},
		},
	})

	ITERATION_CYCLE(model.Event{
		Key:      literals.COMMANDS.DECLARE,
		KeyWord:  q["KeyWord"],
		Payload:  q["Payload"],
		DateTime: low.CurrentTime(),
	})
}
