package field

import (
	"fmt"

	"main/constants"
	"main/model"
)

var field = model.FieldType{
	Position:      0,
	PositionRange: 0,
	Message:       "",
}

func reset() {
	field.Message = ""
	field.Position = 0
	field.PositionRange = 0
}

func setMessage(message string) {
	field.Message = message
}

func incrementPosition() {
	if field.Position < field.PositionRange {
		field.Position++
	}
}

func decrementPosition() {
	if field.Position > 0 {
		field.Position--
	}
}

func incrementPositionRange() {
	field.PositionRange++
}

func decrementPositionRange() {
	if field.PositionRange > 0 {
		field.PositionRange--
	}
}

func horizontalCursorToLeft() {
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
	if position < 0 || position > field.PositionRange {
		return
	}
	fmt.Printf("\033[%dG", len(constants.FIELDPREFIX)+position+2)
	field.Position = position
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
			decrementPositionRange()
		}
		decrementPosition()
	}
}

func horizontalJumpToEnd() {
	field.Position = field.PositionRange
	horizontalCursorToPosition(field.PositionRange)
}
