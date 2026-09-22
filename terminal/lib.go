package terminal

import (
	"fmt"
	"strconv"
	"strings"

	"atomicgo.dev/cursor"
	"github.com/nathan-fiscaletti/consolesize-go"
)

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
	Output("\n")
	cursor.StartOfLine()
}

func GetCustomMessage(message string, SGR ...int) string {
	if len(SGR) == 0 {
		return message
	}
	parts := make([]string, len(SGR))
	for i, v := range SGR {
		parts[i] = strconv.Itoa(v)
	}
	return fmt.Sprintf("\033[%sm%s\033[0m", strings.Join(parts, ";"), message)
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
