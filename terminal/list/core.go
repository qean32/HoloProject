package list

import (
	"main/model"
	"main/terminal"
)

var list = model.List{
	Position: 0,
	Length:   0,
	Options:  []model.Option{},
}

func incrementPosition() bool {
	if list.Position < list.Length-1 {
		list.Position++
		terminal.VerticalJumpDown()
		return true
	}
	return false
}

func decrimentPosition() bool {
	if list.Position > 0 {
		list.Position--
		terminal.VerticalJumpUp()
		return true
	}
	return false
}

func reset() {
	list.Options = []model.Option{}
	list.Position = 0
	list.Length = 0
}

func set(options []model.Option) {
	list.Options = options
	list.Position = 0
	list.Length = len(options)
}
