package terminal

import "fmt"

func VerticalJumpUp() {
	fmt.Print("\033[1A")
}

func VerticalJumpDown() {
	fmt.Print("\033[1B")
}

func VerticalJumpUpCount(count int) {
	fmt.Printf("\033[%dA", count)
}

func VerticalJumpDownCount(count int) {
	fmt.Printf("\033[%dB", count)
}

func HorizontalJumpToStart() {
	fmt.Print("\r")
}

func OutputTechInfo(messages ...any) {
	var result string
	for _, msg := range messages[1:] {
		result += fmt.Sprint(msg)
	}
	ReRenderLine(fmt.Sprint(messages[0]) + "\033[31m Технический вывод: " + result + "\033[0m")
}

func ClearLine() {
	fmt.Print("\r \033[2K")
}

func ClearLines(count int) {
	for i := 0; i < count; i++ {
		ClearLine()
		VerticalJumpUp()
	}
}
