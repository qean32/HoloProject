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

func _setOptions(options []model.Option) {
	list.Options = options
	list.Position = 0
	list.Length = len(options)
}

func IncrementPosition() bool {
	if list.Position < list.Length-1 {
		list.Position++
		terminal.JumpToDown()
		return true
	}
	return false
}

func DecrimentPosition() bool {
	if list.Position > 0 {
		list.Position--
		terminal.JumpToUp()
		return true
	}
	return false
}

func ResetList() {
	list.Options = []model.Option{}
	list.Position = 0
	list.Length = 0
}
