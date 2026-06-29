package terminal

import "fmt"

func MoveCursorLeft() {
	if field.Position > 0 {
		fmt.Print("\033[1D")
		ChangePositionCursor(-1, false)
	}
}

func MoveCursorRight() {
	if field.Position < field.PositionRange {
		fmt.Print("\033[1C")
		ChangePositionCursor(1, false)
	}
}

func MoveCursorToPosition(position int) {
	if position <= field.PositionRange {
		fmt.Printf("\033[%dG", position)
		field.Position = position
	}
}

func ChangePositionCursor(operation int, moveRange bool) {
	if operation != -1 && operation != 1 {
		return
	}

	if operation == 1 {
		if moveRange {
			IncrementPositionRange()
		}
		IncrementPosition()
	} else {
		if moveRange {
			DecrimentPositionRange()
		}
		DecrimentPosition()
	}
}
