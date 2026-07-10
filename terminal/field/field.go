package field

import (
	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"main/lib"
)

func Field() string {
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		char := key.String()

		switch key.Code {
		case keys.Space:
			char = " "
		case keys.Backspace:
			removeChar()
		case keys.Enter:
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
			lib.StopProcess()
		case keys.CtrlC:
			lib.StopProcess()
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
