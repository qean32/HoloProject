package deep

import (
	"fmt"
	"main/constants"
	"strings"

	"github.com/nathan-fiscaletti/consolesize-go"
)

func CONSOLE(output string) {
	fmt.Println(output)
	fmt.Print("~ ")
}

func CONSOLE_CENTER(output string) {
	CONSOLE(strings.Repeat(" ", CalcCenterCMD(len(output))) + output)
}

func CalcCenterCMD(length int) int {
	cols, _ := consolesize.GetConsoleSize()
	return cols/2 - (length / 2)
}

func CONSOLE_ASCII_CENTER(ASCII string) {
	array := strings.Split(ASCII, constants.NEXT_LINE)
	s := strings.Repeat(" ", CalcCenterCMD(len(array[1])-2))
	n := (len(array) - 2) * 2

	params := make([]interface{}, n)
	for i := range params {
		params[i] = s
	}

	CONSOLE(fmt.Sprintf(ASCII, params...))
}

func CONSOLE_RESPONSE(output string, next bool) {
	fmt.Println(constants.OUTPUT_MESSAGE + output)
	if next {
		fmt.Print("~ ")
	}
}
