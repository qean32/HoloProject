package field

import (
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"main/lib/death"
	"main/lib/utils"
	"main/model"
	"main/terminal"
)

func Field(payload model.FieldPayload) string {
	reset()
	set(payload)

	terminal.Print(utils.GetCustomMessage(field.Prefix, field.PrefixSRG...))

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		symbol := key.String()

		switch key.Code {
		case keys.Space:
			symbol = " "
		case keys.Backspace:
			removeSymbol()
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
			death.Exit()
			return true, nil
		case keys.CtrlC:
			death.Exit()
			return true, nil
		}

		if len(symbol) == 1 {
			pushSymbol(symbol)
		}
		return false, nil
	})

	return field.Message
}
