package terminal

import (
	"fmt"
	"main/model"
	"strings"

	"main/lib"

	"atomicgo.dev/cursor"
)

var field = model.FIELDTYPE{
	Position:      0,
	PositionRange: 0,
	Field:         "",
}

func AddChar(char string) {
	OutputTechInfo(field.Field, field)
	if field.Position != len(field.Field) {
		chars := strings.Split(field.Field, "")
		field.Field = strings.Join(lib.AddAfterIndex(chars, char, field.Position), "")

		cursor.ClearLine()
		fmt.Print(field.Field)
		MoveCursorToPosition(field.Position + 1)

		field.PositionRange++
	} else {
		field.Field = field.Field + char
		cursor.ClearLine()
		fmt.Print(field.Field)
		ChangePositionCursor(1, true)
	}
}

func ChangePositionCursor(operation int, moveRange bool) {
	if operation != -1 && operation != 1 {
		return
	}
	if operation == 1 {
		field.Position++
		if moveRange {
			field.PositionRange++
		}
	} else if field.Position != 0 {
		field.Position--
		if moveRange {
			field.PositionRange--
		}
	}
}

func AddSpace() {
	field.Field = field.Field + " "
	ChangePositionCursor(1, true)
	fmt.Print(" ")
}

func RemoveChar() {
	if len(field.Field) != 0 {
		OutputTechInfo(field.Field, field)
		field.Field = strings.Join(lib.RemoveByIndex(strings.Split(field.Field, ""), field.Position), "")
		cursor.Left(1)
		// fmt.Print("\033[K\r")
		MoveCursorToPosition(field.Position)
		fmt.Print(field.Field)
		// 	field.Field = field.Field[:len(field.Field)-1]
	}
}

func MoveCursorLeft() {
	cursor.Left(1)
	if field.Position > 0 {
		ChangePositionCursor(-1, false)
	}
}

func MoveCursorRight() {
	if field.Position < field.PositionRange {
		cursor.Right(1)
		ChangePositionCursor(1, false)
	}
}

func MoveCursorToPosition(position int) {
	if position <= field.PositionRange {
		cursor.HorizontalAbsolute(position)
		field.Position = position
	}
}

func ResetField() {
	field.Field = ""
	field.Position = 0
	field.PositionRange = 0
}
