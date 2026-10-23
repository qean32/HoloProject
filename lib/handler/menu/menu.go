package menu

import (
	eventlist "main/constants/event-list"
	"path/filepath"

	"main/constants"
	errorlist "main/constants/error-list"
	"main/constants/literal"
	"main/lib/array"
	"main/lib/death"
	"main/lib/utils"
	"main/model"
	"main/terminal/menu"
)

var menuOption = []model.Option{
	{Message: "Запустить команду", Event: model.Event{Key: eventlist.BATLIST}},
	{Message: "Ручной ввод", Event: model.Event{Key: eventlist.MANUAL}},
	{Message: "Удалить команду", Event: model.Event{Key: eventlist.BATLISTREMOVE}},
	{Message: "Добавить команду", Event: model.Event{Key: eventlist.MANUCREATEBAT}},
	{Message: "Генерация мастер ключа", Event: model.Event{Key: eventlist.INWORK}},
	{Message: "Помощь", Event: model.Event{Key: eventlist.HELP}},
	{Message: "Выход", Event: model.Event{Key: eventlist.STOP}},
}

func LocalList(title string, makeEvent func(name string) model.Event) error {
	dir := filepath.Join(constants.Root, constants.Cmd)
	commands, err := death.ListFilesByExt(dir, literal.Extension.Bat)
	if err != nil {
		return errorlist.Read(dir, err)
	}

	menu.Menu(model.ListPayload{
		Options: array.Map(commands, func(value string) model.Option {
			return model.Option{
				Message: utils.TrimBatExtension(value),
				Event:   makeEvent(value),
			}
		}),
		Title: title,
	})
	return nil
}

func runBatList() error {
	return LocalList(literal.Titles.ListBat, func(value string) model.Event {
		return model.Event{
			Key:     eventlist.RUNBAT,
			KeyWord: utils.TrimBatExtension(value),
		}
	})
}

func removeBatList() error {
	return LocalList(literal.Titles.ListBat, func(value string) model.Event {
		trimmed := utils.TrimBatExtension(value)
		return model.Event{
			Key:     eventlist.REMOVEBAT,
			KeyWord: trimmed,
			SubEvent: &model.Event{
				Key:     eventlist.PUREREMOVEBAT,
				KeyWord: trimmed,
			},
		}
	})
}
