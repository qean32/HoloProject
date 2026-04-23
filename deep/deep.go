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

var CALLSTACK = []model.Event{}
var TMP_DATA = [][]string{}
var TMP_COMMANDS = [][]string{}

func GenerateMaster() {
}

func NewTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func SET_DATA() {
	if len(TMP_COMMANDS) == 0 {
		strs := (ReadFile(constants.COMMAND_PATH))
		var commands [][]string

		for i := 0; i < len(strs); i++ {
			tmp := strings.Split(strs[i], " ")
			commands = append(commands, []string{tmp[0], strings.Join(tmp[1:], " ")})
		}
		TMP_COMMANDS = commands
	}
	if len(TMP_DATA) == 0 {
		strs := (ReadFile(constants.DATA_PATH))
		var data [][]string

		for i := 0; i < len(strs); i++ {
			tmp := strings.Split(strs[i], " ")
			data = append(data, []string{tmp[0], strings.Join(tmp[1:], " ")})
		}
		TMP_DATA = data
	}
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
			Console(constants.STOP_COMMAND)
		}
	}
}

func RunCMD(command string) {
	cmd := exec.Command("CMD.exe", "/C", command)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при запуске команды: ", err)
		return
	}
}

func Console(output string) {
	fmt.Println(output)
	fmt.Print("$ ")
}
