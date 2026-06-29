package terminal

import (
	"main/lib"
	"strings"
)

func AddChar(char string) {
	if field.Position != len(field.Message) {
		chars := strings.Split(field.Message, "")
		ChangeMessage(strings.Join(lib.AddAfterIndex(chars, char, field.Position), ""))
		RefreshLine(field.Message)
		IncrementPositionRange()
		MoveCursorToPosition(field.Position + 1)
	} else {
		ChangeMessage(field.Message + char)
		RefreshLine(field.Message)
		ChangePositionCursor(1, true)
	}
}

func RemoveChar() {
	if len(field.Message) != 0 {
		DecrimentPosition()
		DecrimentPositionRange()
		ChangeMessage(strings.Join(lib.RemoveByIndex(strings.Split(field.Message, ""), field.Position), ""))
		RefreshLine(field.Message)
		MoveCursorToPosition(field.Position)
	}
}
