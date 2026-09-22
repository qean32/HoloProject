package field

import (
	"main/lib/array"
	"main/terminal"
	"strings"
)

func localReRenderLine(message string) {
	terminal.ReRenderLine(prefix + message)
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
	decrimentPosition()
	decrimentPositionRange()
	setMessage(strings.Join(array.RemoveByIndex(strings.Split(field.Message, ""), field.Position), ""))
	localReRenderLine(field.Message)
	horizontalCursorToPosition(field.Position)
}
