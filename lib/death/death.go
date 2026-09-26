package death

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	"time"

	"main/callstack/event"
	"main/constants/literals"
	"main/model"
)

var reader = bufio.NewReader(os.Stdin)

func Logger(e model.Event) {
	noLog := slices.ContainsFunc(e.Flags, func(item string) bool {
		return strings.TrimSpace(item) == literals.FLAGS.NOLOG
	})
	if !noLog {
		PushToFile(literals.Path.PathLog, fmt.Sprintf("%#v", e))
	}
}

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func RunCmd(command string) {
	cmd := exec.Command("CMD.exe", "/C", command)
	if err := cmd.Start(); err != nil {
		event.Response("$ Ошибка при запуске команды", err)
		return
	}
	if err := cmd.Wait(); err != nil {
		event.Response("$ Команда завершилась с ошибкой", err)
	}
}

func Exit() {
	clearTerminal()
	os.Exit(0)
}

func GetShortEvent(event model.Event) model.Event {
	event.DateTime = CurrentTime()
	return event
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}
