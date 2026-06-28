package terminal

import (
	"fmt"
	"os"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func Field() {
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		char := key.String()

		if len(char) == 1 {
			AddChar(char)
		}
		if key.Code == keys.Space {
			AddSpace()
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
			StopProccess()
		}
		return false, nil
	})
}

func EnterCommand() {
	OutputWithStartLeft("Обработка")
	ResetField()
}

func OutputTechInfo(messages ...any) {
	var result string
	for _, msg := range messages[1:] {
		result += fmt.Sprint(msg)
	}
	OutputWithStartLeft(fmt.Sprint(messages[0]) + "\033[31m Технический вывод: " + result + "\033[0m")
}

func OutputWithStartLeft(message string) {
	cursor.StartOfLineDown(0)
	fmt.Print(message + "\n")
	cursor.StartOfLineDown(0)
}

func StopProccess() {
	os.Exit(0)
}
