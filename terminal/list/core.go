package list

import (
	"main/model"

	"atomicgo.dev/cursor"
)

var list = model.List{
	Position: 0,
	Length:   0,
	Options:  []model.Option{},
}

func incrementPosition() bool {
	if list.Position < list.Length-1 {
		list.Position++
		cursor.Down(1)
		return true
	}
	if list.Length > 1 {
		list.Position = 0
		jumpToStartList()
		cursor.Down(1)
		return true
	}
	return false
}

func decrementPosition() bool {
	if list.Position > 0 {
		list.Position--
		cursor.Up(1)
		return true
	}
	if list.Length > 1 {
		diff := list.Length - 1 - list.Position
		if diff > 0 {
			cursor.Down(diff)
		}
		list.Position = list.Length - 1
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
