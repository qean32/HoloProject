package deep

import (
	"fmt"
	"main/constants"
	"main/model"
	"os/exec"
	"slices"
	"strings"
	"time"
)

func LOG(e model.Event) {
	if slices.IndexFunc(e.Flags, func(item string) bool { return strings.TrimSpace(item) == constants.FLAGS.NOLOG }) == -1 {
		PushToFile(constants.PATH_LOG, fmt.Sprintf("%#v", e))
	}
}

func GenerateMaster() {
}

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ACCESS_ACTION() bool {
	var response string
	fmt.Print("need access action (yes|no) ~ ")
	fmt.Scan(&response)

	if slices.Contains(constants.ACCESS_VARIANTS, response) {
		return true
	}

	return false
}

func DECORATOR_ACCESS_ACTION(function model.EventFunction) model.EventFunction {
	return func(e model.Event) {

		if ACCESS_ACTION() {
			function(e)
		} else {
			CONSOLE_RESPONSE(constants.STOP_COMMAND, false)
		}
	}
}

func RUN_CMD(command string) {
	cmd := exec.Command("CMD.exe", "/C", command)
	err := cmd.Run()
	if err != nil {
		fmt.Println("$ Ошибка при запуске команды", err)
		return
	}
}
