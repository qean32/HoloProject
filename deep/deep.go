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
	if slices.IndexFunc(e.Flags, func(item string) bool {
		return strings.TrimSpace(item) == "-nl"
	}) == -1 {
		PushToFile(constants.LOG_PATH, fmt.Sprintf("%#v", e))
	}
}

func GenerateMaster() {
}

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ACCESS_ACTION() bool {
	var response string
	fmt.Print("Need access (yes/no): ")
	fmt.Scan(&response)

	if response == "yes" || response == "y" || response == "yea" {
		return true
	}

	return false
}

func DECORATOR_ACCESS_ACTION(f model.EventFunction) model.EventFunction {
	return func(e model.Event) {

		if ACCESS_ACTION() {
			f(e)
		} else {
			CONSOLE(constants.STOP_COMMAND)
		}
	}
}

func RUN_CMD(command string) {
	cmd := exec.Command("CMD.exe", "/C", command)
	err := cmd.Run()
	if err != nil {
		fmt.Println("$ Ошибка при запуске команды: ", err)
		return
	}
}

func CONSOLE(output string) {
	fmt.Println(output)
	fmt.Print("$ ")
}
