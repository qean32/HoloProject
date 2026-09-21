package terminal

import (
	"fmt"
	"main/lib/array"
	"strconv"
	"strings"

	"atomicgo.dev/cursor"
	"github.com/nathan-fiscaletti/consolesize-go"
)

func OutputTechInfo(messages ...any) {
	var result string
	for _, msg := range messages {
		result += fmt.Sprint(msg)
	}
	Outputln("\033[31m Технический вывод: " + result + "\033[0m")
}

func OutputCenter(output string, separator string) {
	Output(strings.Repeat(separator, CalcCenterCMD(len(output))) + output)
}

func OutputASCII_CENTER(ASCII string, separator string) {
	lines := strings.Split(ASCII, "\n")
	if len(lines) < 2 {
		Output(ASCII)
		return
	}

	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	offset := CalcCenterCMD(maxLen - 2)
	if offset < 0 {
		offset = 0
	}

	repeat := strings.Repeat(separator, offset)
	length := len(lines) - 2

	params := make([]interface{}, length)
	for i := range params {
		params[i] = repeat
	}

	Output(fmt.Sprintf(ASCII, params...))
}

func CalcCenterCMD(length int) int {
	cols, _ := consolesize.GetConsoleSize()
	res := cols/2 - (length / 2)
	if res < 0 {
		return 0
	}
	return res
}

func DownAndStart() {
	fmt.Print("\n")
	cursor.StartOfLine()
}

func ClearLines(count int) {
	for i := 0; i < count; i++ {
		cursor.ClearLine()
		cursor.Up(1)
	}
}

func ReRenderLine(_message string) {
	cursor.ClearLine()
	cursor.StartOfLine()
	Output(_message)
}

func GetCustomMessage(message string, SGR ...int) string {
	if len(SGR) == 0 {
		return message
	}

	SGRstring := strings.Join(
		array.Map(SGR, func(value int) string {
			return strconv.Itoa(value)
		}), ";")

	return fmt.Sprintf("\033[%sm%s\033[0m", SGRstring, message)
}

func Output(message string) {
	fmt.Print(message)
}

func Outputln(message string) {
	fmt.Println(message)
}
