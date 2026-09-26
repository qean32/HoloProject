package terminal

import (
	"fmt"
	"strconv"
	"strings"

	"atomicgo.dev/cursor"
	"github.com/nathan-fiscaletti/consolesize-go"
)

func PrintCenter(output string, separator string) {
	Print(strings.Repeat(separator, calcCenterCMD(len(output))) + output)
}

func PrintASCIICenter(ascii string, separator string) {
	lines := strings.Split(ascii, "\n")
	if len(lines) < 2 {
		Print(ascii)
		return
	}

	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	offset := calcCenterCMD(maxLen - 2)
	if offset < 0 {
		offset = 0
	}

	repeat := strings.Repeat(separator, offset)
	length := len(lines) - 2

	params := make([]any, length)
	for i := range params {
		params[i] = repeat
	}

	Print(fmt.Sprintf(ascii, params...))
}

func calcCenterCMD(length int) int {
	cols, _ := consolesize.GetConsoleSize()
	res := cols/2 - length/2
	if res < 0 {
		return 0
	}
	return res
}

func DownAndStart() {
	Println()
	cursor.StartOfLine()
}

func GetCustomMessage(message string, sgr ...int) string {
	if len(sgr) == 0 {
		return message
	}
	parts := make([]string, len(sgr))
	for i, v := range sgr {
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

func ReRenderLine(message string) {
	cursor.StartOfLine()
	cursor.ClearLine()
	Print(message)
}
