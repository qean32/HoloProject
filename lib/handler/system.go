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
	"main/terminal"
)

const dirPerm os.FileMode = 0o755

func reset() {
	death.ClearTerminal()
	terminal.RenderBaner()
	manual.Manual()
}

func drop() error {
	if err := os.RemoveAll(constants.Root); err != nil {
		return fmt.Errorf("Удаление %s: %w", constants.Root, err)
	}
	if err := os.MkdirAll(constants.Root, dirPerm); err != nil {
		return fmt.Errorf("Создание %s: %w", constants.Root, err)
	}
	return nil
}

func clearLog() error {
	if err := death.ClearFile(literal.Path.PathLog); err != nil {
		return fmt.Errorf("очистка лога: %w", err)
	}
	return nil
}

func openLog() error {
	path := filepath.Join(constants.Root, literal.Path.PathLog)
	if err := death.RunCmd(path); err != nil {
		return fmt.Errorf("открытие %s: %w", path, err)
	}
	return nil
}

func inwork() {
	event.Response("В разработке")
}

func init_() {
	os.Mkdir(constants.Root, 0755)
	os.Mkdir(filepath.Join(constants.Root, constants.Cmd), 0755)
	death.CreateFile(literal.Path.PathLog, "")
}
