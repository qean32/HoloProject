package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"main/constants"
	eventlist "main/constants/event-list"
	"main/constants/literal"
	"main/lib/array"
	"main/lib/death"
	"main/model"
	"main/terminal"
	"main/terminal/list"
	"main/terminal/questionnaire"
)

func _response(event model.Event) error {
	terminal.PrintResponse(event.Payload)
	return nil
}

func bat(event model.Event) error {
	if err := validateName(event.KeyWord); err != nil {
		return fmt.Errorf("Некорректное имя %q: %w", event.KeyWord, err)
	}
	if err := death.CreateFile(cmdPath(event.KeyWord), event.Payload); err != nil {
		return fmt.Errorf("Создание %s: %w", event.KeyWord, err)
	}
	return nil
}

func runBat(event model.Event) error {
	if err := validateName(event.KeyWord); err != nil {
		return fmt.Errorf("Некорректное имя %q: %w", event.KeyWord, err)
	}
	if err := death.RunCmd(cmdPath(event.KeyWord)); err != nil {
		return fmt.Errorf("Запуск %s: %w", event.KeyWord, err)
	}
	return nil
}

func removeCmd(event model.Event) error {
	if err := validateName(event.KeyWord); err != nil {
		return fmt.Errorf("Некорректное имя %q: %w", event.KeyWord, err)
	}
	if err := os.Remove(cmdPath(event.KeyWord)); err != nil {
		return fmt.Errorf("Удаление %s: %w", event.KeyWord, err)
	}
	return nil
}

func listBatFiles(title string, makeEvent func(name string) model.Event) error {
	dir := filepath.Join(constants.Root, constants.Cmd)
	commands, err := death.ListFilesByExt(dir, literal.Extension.Bat)
	if err != nil {
		return fmt.Errorf("Чтение %s: %w", dir, err)
	}

	list.List(model.ListPayload{
		Options: array.Map(commands, func(v string) model.Option {
			return model.Option{
				Message: strings.TrimSuffix(v, literal.Extension.Bat),
				Event:   makeEvent(v),
			}
		}),
		Title: title,
	})
	return nil
}

func batList(event model.Event) error {
	return listBatFiles(literal.Titles.ListBat, func(value string) model.Event {
		return model.Event{Key: event.Key, KeyWord: value}
	})
}

func runBatList(_ model.Event) error {
	return listBatFiles(literal.Titles.ListBat, func(value string) model.Event {
		return model.Event{
			Key:     eventlist.RUNBAT,
			KeyWord: trimExtention(value),
		}
	})
}

func removeBatList(_ model.Event) error {
	return listBatFiles(literal.Titles.ListBat, func(value string) model.Event {
		trimmed := trimExtention(value)
		return model.Event{
			Key:     eventlist.RRBAT,
			KeyWord: trimmed,
			SubEvent: &model.Event{
				Key:     eventlist.RRBAT,
				KeyWord: trimmed,
			},
		}
	})
}

func questionnaireAddBat(_ model.Event) error {
	q := questionnaire.Questionnaire(model.QuestionnairePayload{
		Questions: []model.Question{
			{
				Message: "Ключ",
				Key:     "KeyWord",
				Callback: func(res string) bool {
					return validateName(res) == nil
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

	return bat(model.Event{
		Key:      eventlist.BAT,
		KeyWord:  q["KeyWord"],
		Payload:  q["Payload"],
		DateTime: death.CurrentTime(),
	})
}
