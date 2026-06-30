package list

import (
	"main/model"
	"main/terminal"
)

var list = model.List{
	Position: 0,
	Options:  map[string]model.EventFunction{},
}

func IncrementPosition() {
	if list.Position < len(list.Options) {
		list.Position++
		terminal.JumpToUp()
	}
}

func DecrimentPosition() {
	if list.Position > 0 {
		list.Position--
		terminal.JumpToDown()
	}
}

func ResetList() {
	list.Options = map[string]model.EventFunction{}
	list.Position = 0
}
