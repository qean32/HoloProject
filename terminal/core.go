package terminal

import (
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

func decrimentPosition() {
	if field.Position > 0 {
		field.Position--
	}
}

func incrementPositionRange() {
	field.PositionRange++
}

func decrimentPositionRange() {
	if field.PositionRange > 0 {
		field.PositionRange--
	}
}
