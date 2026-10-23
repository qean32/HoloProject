package handler

import (
	"os"
	"os/exec"
	"path/filepath"

	"main/callstack/event"
	"main/constants"
	errorlist "main/constants/error-list"
	"main/constants/literal"
	"main/lib/death"
	"main/lib/eventlog"
	"main/terminal"
)

func clearScreen() {
	_ = death.ClearTerminal()
	terminal.RenderBanner()
}

func drop() error {
	if !filepath.IsAbs(constants.Root) {
		return errorlist.ErrForbiddenName
	}
	if err := os.RemoveAll(constants.Root); err != nil {
		return errorlist.Remove(constants.Root, err)
	}
	// if err := constants.EnsureLayout(); err != nil {
	// 	return err
	// }
	return death.EnsureFile(literal.Path.PathLog)
}

func clearLog() error {
	if err := death.ClearFile(literal.Path.PathLog); err != nil {
		return errorlist.ClearLog(err)
	}
	return nil
}

func openLog() error {
	path := filepath.Join(constants.Root, literal.Path.PathLog)
	if _, err := os.Stat(path); err != nil {
		return errorlist.Open(path, err)
	}
	if err := exec.Command("cmd", "/c", "start", "", path).Run(); err != nil {
		return errorlist.Open(path, err)
	}
	return nil
}

func inwork() {
	event.Response("В разработке")
}

func toggleLogger() {
	if eventlog.Toggle() {
		event.Response("Логирование событий: включено")
		return
	}
	event.Response("Логирование событий: выключено")
}

func init_() error {
	// if err := constants.EnsureLayout(); err != nil {
	// 	return err
	// }
	return death.EnsureFile(literal.Path.PathLog)
}
