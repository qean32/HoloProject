package flow

import (
	"path/filepath"
	"strings"

	"main/constants"
	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
	"main/constants/literal"
	"main/lib/array"
	"main/lib/death"
	names "main/lib/names"
	"main/model"
	"main/terminal/list"
	"main/terminal/questionnaire"
)

var menuOptions = []model.Option{
	{Message: "Запустить команду", Event: model.Event{Key: eventlist.BATLIST}},
	{Message: "Удалить команду", Event: model.Event{Key: eventlist.BATLISTREMOVE}},
	{Message: "Добавить команду", Event: model.Event{Key: eventlist.MANUCREATEBAT}},
	{Message: "Очистить экран", Event: model.Event{Key: eventlist.CLEAR}},
	{Message: "Помощь", Event: model.Event{Key: eventlist.HELP}},
	{Message: "Выход", Event: model.Event{Key: eventlist.STOP}},
}

// Menu показывает главное меню и возвращает выбранный пункт.
func Menu(_ model.Event) (model.Action, error) {
	return model.ActionFunc(func() ([]model.Event, error) {
		selected, ok := list.List(model.ListPayload{
			Options: menuOptions,
			Title:   literal.Titles.Menu,
		})
		if !ok {
			return nil, nil
		}
		return []model.Event{selected}, nil
	}), nil
}

// CreateBat спрашивает имя и содержимое .bat и возвращает событие BAT.
func CreateBat(_ model.Event) (model.Action, error) {
	return model.ActionFunc(func() ([]model.Event, error) {
		answers := questionnaire.Questionnaire(model.QuestionnairePayload{
			Questions: []model.Question{
				{
					Message: "Ключ",
					Key:     "KeyWord",
					Callback: func(res string) bool {
						return names.Validate(res) == nil
					},
				},
				{
					Message: "Команда",
					Key:     "Payload",
					Callback: func(res string) bool {
						return strings.TrimSpace(res) != ""
					},
				},
			},
			Title: literal.Titles.EnterBat,
		})

		return []model.Event{{
			Key:      eventlist.BAT,
			KeyWord:  answers["KeyWord"],
			Payload:  answers["Payload"],
			DateTime: death.CurrentTime(),
		}}, nil
	}), nil
}

// ConfirmRemove спрашивает подтверждение и возвращает вложенное событие удаления.
func ConfirmRemove(event model.Event) (model.Action, error) {
	if event.SubEvent == nil {
		return nil, errorlist.ErrNoSubEvent
	}
	next := *event.SubEvent

	return model.ActionFunc(func() ([]model.Event, error) {
		answers:= questionnaire.Questionnaire(model.QuestionnairePayload{
			Questions: []model.Question{
				{
					Message:  "Подтвердите действие (yea/no)",
					Key:      "Response",
					Callback: func(string) bool { return true },
				},
			},
		})
		if !questionnaire.IsAcceptableYea(answers["Response"]) {
			return nil, errorlist.ErrAccessDenied
		}
		return []model.Event{next}, nil
	}), nil
}

// RunBatList показывает список .bat и возвращает событие RUNBAT.
func RunBatList(_ model.Event) (model.Action, error) {
	return selectBat(func(name string) model.Event {
		return model.Event{
			Key:     eventlist.RUNBAT,
			KeyWord: name,
		}
	})
}

// RemoveBatList показывает список .bat и возвращает событие REMOVEBAT.
func RemoveBatList(_ model.Event) (model.Action, error) {
	return selectBat(func(name string) model.Event {
		return model.Event{
			Key:     eventlist.REMOVEBAT,
			KeyWord: name,
			SubEvent: &model.Event{
				Key:     eventlist.PUREREMOVEBAT,
				KeyWord: name,
			},
		}
	})
}

func selectBat(makeEvent func(name string) model.Event) (model.Action, error) {
	return model.ActionFunc(func() ([]model.Event, error) {
		dir := filepath.Join(constants.Root, constants.Cmd)
		files, err := death.ListFilesByExt(dir, literal.Extension.Bat)
		if err != nil {
			return nil, errorlist.Read(dir, err)
		}

		options := array.Map(files, func(file string) model.Option {
			name := strings.TrimSuffix(file, literal.Extension.Bat)
			return model.Option{
				Message: name,
				Event:   makeEvent(name),
			}
		})

		selected, ok := list.List(model.ListPayload{
			Options: options,
			Title:   literal.Titles.ListBat,
		})
		if !ok {
			return nil, nil
		}
		return []model.Event{selected}, nil
	}), nil
}
