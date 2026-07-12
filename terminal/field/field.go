package field

import (
	"main/lib/low"
	"main/terminal"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func Field() string {
	reset()
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		char := key.String()

		switch key.Code {
		case keys.Space:
			char = " "
		case keys.Backspace:
			removeChar()
		case keys.Enter:
			terminal.DownAndStart()
			return true, nil
		case keys.Left:
			horizontalCursorToLeft()
		case keys.Right:
			horizontalCursorToRight()
		case keys.End:
			horizontalJumpToEnd()
		case keys.Home:
			cursor.StartOfLine()
		case keys.Escape:
			low.StopProcess()
		case keys.CtrlC:
			low.StopProcess()
		}

		if len(char) == 1 {
			pushChar(char)
		}
		return false, nil
	})

	return field.Message
}

func enter() {
	reset()
}
