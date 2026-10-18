package handler

import (
	"os"
	"path/filepath"

	"main/callstack/event"
	"main/callstack/manual"
	"main/constants"
	errorlist "main/constants/error-list"
	"main/constants/literal"
	"main/lib/death"
	"main/terminal"
)

const dirPerm os.FileMode = 0o755

func reset() {
	death.ClearTerminal()
	terminal.RenderBanner()
	manual.Manual()
}

func drop() error {
	if err := os.RemoveAll(constants.Root); err != nil {
		return errorlist.Remove(constants.Root, err)
	}
	if err := os.MkdirAll(constants.Root, dirPerm); err != nil {
		return errorlist.Create(constants.Root, err)
	}
	return nil
}

func clearLog() error {
	if err := death.ClearFile(literal.Path.PathLog); err != nil {
		return errorlist.ClearLog(err)
	}
	return nil
}

func openLog() error {
	path := filepath.Join(constants.Root, literal.Path.PathLog)
	if err := death.RunCmd(path); err != nil {
		return errorlist.Open(path, err)
	}
	return nil
}

func inwork() {
	event.Response("В разработке")
}

func init_() error {
	if err := os.MkdirAll(filepath.Join(constants.Root, constants.Cmd), dirPerm); err != nil {
		return errorlist.Create(constants.Cmd, err)
	}
	return death.CreateFile(literal.Path.PathLog, "")
}
