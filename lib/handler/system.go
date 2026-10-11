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
	terminal.RenderBaner()
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

func init_() {
	os.Mkdir(constants.Root, 0755)
	os.Mkdir(filepath.Join(constants.Root, constants.Cmd), 0755)
	death.CreateFile(literal.Path.PathLog, "")
}
