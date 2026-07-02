package terminal

import (
	"main/lib"
	"strings"
)

func pushChar(char string) {
	if field.Position != len(field.Message) {
		chars := strings.Split(field.Message, "")
		setMessage(strings.Join(lib.AddAfterIndex(chars, char, field.Position), ""))
		ReRenderLine(field.Message)
		incrementPositionRange()
		horizontalCursorToPosition(field.Position + 1)
	} else {
		setMessage(field.Message + char)
		ReRenderLine(field.Message)
		changePositionCursor(1, true)
	}
}

func removeChar() {
	if len(field.Message) != 0 {
		decrimentPosition()
		decrimentPositionRange()
		setMessage(strings.Join(lib.RemoveByIndex(strings.Split(field.Message, ""), field.Position), ""))
		ReRenderLine(field.Message)
		horizontalCursorToPosition(field.Position)
	}
}
