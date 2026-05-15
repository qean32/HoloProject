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

func CONSOLE_CENTER(output string, separator string) {
	CONSOLE(strings.Repeat(separator, CalcCenterCMD(len(output))) + output)
}

func CalcCenterCMD(length int) int {
	cols, _ := consolesize.GetConsoleSize()
	return cols/2 - (length / 2)
}

func CONSOLE_ASCII_CENTER(ASCII string, separator string) {
	array := strings.Split(ASCII, constants.NEXT_LINE)
	repeat := strings.Repeat(separator, CalcCenterCMD(len(array[1])-2))
	length := len(array) - 2

	params := make([]interface{}, length)
	for i := range params {
		params[i] = repeat
	}

	CONSOLE(fmt.Sprintf(ASCII, params...))
}

func CONSOLE_RESPONSE(output string, next bool) {
	fmt.Println(constants.OUTPUT_MESSAGE + output)
	if next {
		fmt.Print("~ ")
	}
}
