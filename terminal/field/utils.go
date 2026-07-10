package field

import (
	"main/lib/array"
	"main/terminal"
	"strings"
)

func pushChar(char string) {
	if field.Position != len(field.Message) {
		chars := strings.Split(field.Message, "")
		setMessage(strings.Join(array.AddAfterIndex(chars, char, field.Position), ""))
		terminal.ReRenderLine(field.Message)
		incrementPositionRange()
		horizontalCursorToPosition(field.Position + 1)
	} else {
		setMessage(field.Message + char)
		terminal.ReRenderLine(field.Message)
		changePositionCursor(1, true)
	}
}

func removeChar() {
	if len(field.Message) != 0 {
		decrimentPosition()
		decrimentPositionRange()
		setMessage(strings.Join(array.RemoveByIndex(strings.Split(field.Message, ""), field.Position), ""))
		terminal.ReRenderLine(field.Message)
		horizontalCursorToPosition(field.Position)
	}
}
