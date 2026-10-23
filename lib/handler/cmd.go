package handler

import (
	"path/filepath"

	"main/constants"
	errorlist "main/constants/error-list"
	"main/lib/death"
	"main/model"
	"main/terminal"
)

func _response(event model.Event) error {
	terminal.PrintResponse(event.Payload)
	return nil
}

func bat(event model.Event) error {
	if err := validateName(event.KeyWord); err != nil {
		return errorlist.InvalidName(event.KeyWord, err)
	}
	if err := death.CreateFile(cmdPath(event.KeyWord), event.Payload); err != nil {
		return errorlist.Create(event.KeyWord, err)
	}
	return nil
}

func runBat(event model.Event) error {
	if err := validateName(event.KeyWord); err != nil {
		return errorlist.InvalidName(event.KeyWord, err)
	}
	if err := death.RunCmd(filepath.Join(constants.Root, cmdPath(event.KeyWord))); err != nil {
		return errorlist.Run(event.KeyWord, err)
	}
	return nil
}

func removeCmd(event model.Event) error {
	if err := validateName(event.KeyWord); err != nil {
		return errorlist.InvalidName(event.KeyWord, err)
	}
	if err := death.RemoveFile(cmdPath(event.KeyWord)); err != nil {
		return errorlist.Remove(event.KeyWord, err)
	}
	return nil
}
