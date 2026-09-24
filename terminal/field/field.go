package field

import (
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"main/constants"
	"main/lib/low"
	"main/terminal"
)

func Field() string {
	terminal.Print(constants.FIELDPREFIX)
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
			horizontalCursorToPosition(0)
		case keys.Escape:
			low.Exit()
			return true, nil
		case keys.CtrlC:
			low.Exit()
			return true, nil
		}

		if len(char) == 1 {
			pushChar(char)
		}
		return false, nil
	})

	return field.Message
}
