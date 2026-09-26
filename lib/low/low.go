package low

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	"time"

	"main/callstack"
	"main/constants"
	"main/constants/literals"
	"main/model"
)

var reader = bufio.NewReader(os.Stdin)

func LOG(e model.Event) {
	noLog := slices.ContainsFunc(e.Flags, func(item string) bool {
		return strings.TrimSpace(item) == literals.FLAGS.NOLOG
	})
	if !noLog {
		PushToFile(constants.PATH_LOG, fmt.Sprintf("%#v", e))
	}
}

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func RUN_CMD(command string) {
	cmd := exec.Command("CMD.exe", "/C", command)
	if err := cmd.Start(); err != nil {
		callstack.PushCallStack(model.Event{
			Key:     literals.COMMANDLIST.RESPONSE,
			Payload: "$ Ошибка при запуске команды",
		})
		return
	}
	if err := cmd.Wait(); err != nil {
		callstack.PushCallStack(model.Event{
			Key:     literals.COMMANDLIST.RESPONSE,
			Payload: "$ Команда завершилась с ошибкой",
		})
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
