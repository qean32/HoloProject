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

		switch key.Code {
		case keys.Space:
			char = " "
		case keys.Backspace:
			removeChar()
		case keys.Enter:
			Enter()
			return true, nil
		case keys.Left:
			horizontalCursorToLeft()
		case keys.Right:
			horizontalCursorToRight()
		case keys.End:
			horizontalJumpToEnd()
		case keys.Home:
			HorizontalJumpToStart()
		case keys.Escape:
			lib.StopProcess()
		case keys.CtrlC:
			lib.StopProcess()
		}

		if len(char) == 1 {
			pushChar(char)
		}
		return false, nil
	})
}

func ReRenderLine(_message string) {
	ClearLine()
	HorizontalJumpToStart()
	fmt.Print(_message)
}

func Enter() {
	ReRenderLine("Обработка")
	reset()
}
