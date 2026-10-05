package handler

import (
	"fmt"
	"os"
	"path/filepath"

	"main/callstack/event"
	"main/callstack/manual"
	"main/constants"
	"main/constants/literal"
	"main/lib/death"
	"main/model"
	"main/terminal"
)

const dirPerm os.FileMode = 0o755

func reset(_ model.Event) error {
	death.ClearTerminal()
	terminal.RenderBaner()
	manual.Manual()
	return nil
}

func drop(_ model.Event) error {
	if err := os.RemoveAll(constants.Root); err != nil {
		return fmt.Errorf("удаление %s: %w", constants.Root, err)
	}
	if err := os.MkdirAll(constants.Root, dirPerm); err != nil {
		return fmt.Errorf("создание %s: %w", constants.Root, err)
	}
	return nil
}

func clearLog(_ model.Event) error {
	if err := death.ClearFile(literal.Path.PathLog); err != nil {
		return fmt.Errorf("очистка лога: %w", err)
	}
	return nil
}

func openLog(_ model.Event) error {
	path := filepath.Join(constants.Root, literal.Path.PathLog)
	if err := death.RunCmd(path); err != nil {
		return fmt.Errorf("открытие %s: %w", path, err)
	}
	return nil
}

func inwork(_ model.Event) error {
	event.Response("В разработке")
	return nil
}
