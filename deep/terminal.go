package deep

// DEPRECETED !!!

import (
	"fmt"
	"main/constants"
	"main/model"
	"slices"
	"strings"

	"github.com/nathan-fiscaletti/consolesize-go"
)

func CONSOLE(output string) {
	fmt.Print(output)
}

func CONSOLE_CENTER(output string, separator string) {
	CONSOLE(strings.Repeat(separator, CalcCenterCMD(len(output))) + output)
}

func CalcCenterCMD(length int) int {
	cols, _ := consolesize.GetConsoleSize()
	res := cols/2 - (length / 2)
	if res < 0 {
		return 0
	}
	return res
}

func CONSOLE_ASCII_CENTER(ASCII string, separator string) {
	array := strings.Split(ASCII, constants.NEXT_LINE)
	offet := CalcCenterCMD(len(array[1]) - 2)

	repeat := strings.Repeat(separator, offet)
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

func ACCESS_ACTION() bool {
	var response string
	fmt.Print("u need access action (yes|no) ~ ")
	fmt.Scanln(&response)
	trimResponse := strings.TrimSpace(response)

	if slices.Contains(constants.ACCESS_VARIANTS, trimResponse) {
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
