package deep

import (
	"bufio"
	"fmt"
	"main/constants"
	"main/constants/literals"
	"main/model"
	"os"
	"os/exec"
	"slices"
	"strings"
	"time"
)

var READER = bufio.NewReader(os.Stdin)

func LOG(e model.Event) {
	if slices.IndexFunc(e.Flags, func(item string) bool { return strings.TrimSpace(item) == literals.FLAGS.NOLOG }) == -1 {
		PushToFile(constants.PATH_LOG, fmt.Sprintf("%#v", e))
	}
}

func GenerateMaster() {
}

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func RUN_CMD(command string) {
	cmd := exec.Command("CMD.exe", "/C", command)
	err := cmd.Run()
	if err != nil {
		fmt.Println("$ Ошибка при запуске команды", err)
		return
	}
}
