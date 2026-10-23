package death

import (
	"os"
	"os/exec"
	"slices"
	"strings"
	"time"

	"main/callstack"
	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
	"main/constants/literal"
	"main/model"
)

func Logger(stackErr model.StackError) error {
	if stackErr.Event.Key == eventlist.RESPONSE {
		return nil
	}
	if noLog := slices.ContainsFunc(stackErr.Event.Flags, func(item string) bool {
		return strings.TrimSpace(item) == literal.Flags.NOLOG
	}); noLog {
		return nil
	}

	safe := stackErr.Event
	safe.Password = ""
	return PushToFile(literal.Path.PathLog, callstack.StackErrorToString(model.StackError{Event: safe, Err: stackErr.Err}, 0))
}

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func RunCmd(command string) error {
	cmd := exec.Command("cmd", "/C", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errorlist.RunCommand(command, err)
	}
	return nil
}

func Exit() {
	_ = ClearTerminal()
	os.Exit(0)
}

func ClearTerminal() error {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return errorlist.ClearTerminal(err)
	}
	return nil
}
