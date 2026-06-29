package terminal

import (
	"fmt"
	"main/model"
)

var field = model.FIELDTYPE{
	Position:      0,
	PositionRange: 0,
	Message:       "",
}

func ResetField() {
	field.Message = ""
	field.Position = 0
	field.PositionRange = 0
}

func IncrementPosition() {
	if field.Position < field.PositionRange {
		field.Position++
	}
}

func DecrimentPosition() {
	if field.Position > 0 {
		field.Position--
	}
}

func IncrementPositionRange() {
	field.PositionRange++
}

func DecrimentPositionRange() {
	if field.PositionRange > 0 {
		field.PositionRange--
	}
}

func ChangeMessage(_message string) {
	field.Message = _message
}

func ClearLine() {
	fmt.Print("\033[2K")
}

func JumpToStartLine() {
	fmt.Print("\r")
}
