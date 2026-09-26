package field

import (
	"strings"

	"main/lib/array"
	"main/lib/utils"
	"main/terminal"
)

func localReRenderLine(message string) {
	terminal.ReRenderLine(utils.GetCustomMessage(field.Prefix+message, field.PrefixSRG...))
}

func pushSymbol(char string) {
	if field.Position != len(field.Message) {
		chars := strings.Split(field.Message, "")
		setMessage(strings.Join(array.AddAfterIndex(chars, char, field.Position), ""))
		localReRenderLine(field.Message)
		incrementPositionRange()
		incrementPosition()
		horizontalCursorToPosition(field.Position)
	} else {
		setMessage(field.Message + char)
		localReRenderLine(field.Message)
		changePositionCursor(1, true)
	}
}

func removeSymbol() {
	if len(field.Message) == 0 {
		return
	}
	decrementPosition()
	decrementPositionRange()
	setMessage(strings.Join(array.RemoveByIndex(strings.Split(field.Message, ""), field.Position), ""))
	localReRenderLine(field.Message)
	horizontalCursorToPosition(field.Position)
}
