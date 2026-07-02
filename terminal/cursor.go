package terminal

import "fmt"

func horizontalCursorToLeft() {
	// DO NOT USE OUTSIDE FIELD !!!
	if field.Position > 0 {
		fmt.Print("\033[1D")
		changePositionCursor(-1, false)
	}
}

func horizontalCursorToRight() {
	if field.Position < field.PositionRange {
		fmt.Print("\033[1C")
		changePositionCursor(1, false)
	}
}

func horizontalCursorToPosition(position int) {
	if position <= field.PositionRange {
		fmt.Printf("\033[%dG", position)
		field.Position = position
	}
}

func changePositionCursor(operation int, moveRange bool) {
	if operation != -1 && operation != 1 {
		return
	}

	if operation == 1 {
		if moveRange {
			incrementPositionRange()
		}
		incrementPosition()
	} else {
		if moveRange {
			decrimentPositionRange()
		}
		decrimentPosition()
	}
}

func horizontalJumpToEnd() {
	field.Position = field.PositionRange
	horizontalCursorToPosition(field.PositionRange)
}
