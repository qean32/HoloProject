package lib

import (
	"bufio"
	"main/constants"
	"main/deep"
	"main/terminal"
	"os"
)

var READER = bufio.NewReader(os.Stdin)

func INIT() {
	terminal.OutputASCII_CENTER(constants.BinaryPROJECT_INIT, " ")
	terminal.OutputASCII_CENTER(constants.PROJECT_INIT, " ")
	constants.INIT_ROOT()
	deep.SETDATA()
}

func Help() {
	terminal.OutputASCII_CENTER(constants.HelpMessage, "")
}

func ClearLog() {
	deep.ClearFile(constants.PATH_LOG)
}

func RunCommand(command string) {
	deep.RUN_CMD(command)
}
