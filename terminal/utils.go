package terminal

import (
	"atomicgo.dev/cursor"
	"fmt"
)

func OutputTechInfo(messages ...any) {
	var result string
	for _, msg := range messages[1:] {
		result += fmt.Sprint(msg)
	}
	ReRenderLine(fmt.Sprint(messages[0]) + "\033[31m Технический вывод: " + result + "\033[0m")
}

func ClearLines(count int) {
	for i := 0; i < count; i++ {
		cursor.ClearLine()
		cursor.Up(1)
	}
}
