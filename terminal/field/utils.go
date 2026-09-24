package field

import (
	"strings"

	"main/constants"
	"main/lib/array"
	"main/terminal"
)

func localReRenderLine(message string) {
	terminal.ReRenderLine(constants.FIELDPREFIX + message)
}

func pushChar(char string) {
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

func removeChar() {
	if len(field.Message) == 0 {
		return
	}
	decrementPosition()
	decrementPositionRange()
	setMessage(strings.Join(array.RemoveByIndex(strings.Split(field.Message, ""), field.Position), ""))
	localReRenderLine(field.Message)
	horizontalCursorToPosition(field.Position)
}
