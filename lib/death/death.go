package death

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	"time"

	eventlist "main/constants/event-list"
	"main/constants/literal"
	"main/model"
)

func Logger(event model.Event) error {
	if event.Key == eventlist.RESPONSE {
		return nil
	}
	noLog := slices.ContainsFunc(event.Flags, func(item string) bool {
		return strings.TrimSpace(item) == literal.Flags.NOLOG
	})
	if noLog {
		return nil
	}
	return PushToFile(literal.Path.PathLog, fmt.Sprintf("%#v", event))
}

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func RunCmd(command string) error {
	cmd := exec.Command("CMD.exe", "/C", command)
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("запуск команды %q: %w", command, err)
	}
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("выполнение команды %q: %w", command, err)
	}
	return nil
}

func Exit() {
	ClearTerminal()
	os.Exit(0)
}

func GetShortEvent(event model.Event) model.Event {
	event.DateTime = CurrentTime()
	return event
}

func ClearTerminal() error {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("очистка терминала: %w", err)
	}
	return nil
}
