package terminal

import (
	"fmt"
	"main/lib"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func Field() {
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		char := key.String()

		if key.Code == keys.Space {
			char = " "
		}
		if len(char) == 1 {
			AddChar(char)
		}
		if key.Code == keys.Backspace {
			RemoveChar()
		}
		if key.Code == keys.Enter {
			EnterCommand()
			return false, nil
		}
		if key.Code == keys.Left {
			MoveCursorLeft()
		}
		if key.Code == keys.Right {
			MoveCursorRight()
		}
		if key.Code == keys.Escape || key.Code == keys.CtrlC {
			lib.StopProcess()
		}
		return false, nil
	})
}

func RefreshLine(_message string) {
	ClearLine()
	JumpToStartLine()
	fmt.Print(_message)
}

func EnterCommand() {
	RefreshLine("Обработка")
	ResetField()
}

func OutputTechInfo(messages ...any) {
	var result string
	for _, msg := range messages[1:] {
		result += fmt.Sprint(msg)
	}
	RefreshLine(fmt.Sprint(messages[0]) + "\033[31m Технический вывод: " + result + "\033[0m")
}
