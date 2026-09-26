package list

import (
	"main/model"
)

var list = model.List{
	Title:    "",
	Position: 0,
	Length:   0,
	Options:  []model.Option{},
}

func reset() {
	list.Options = []model.Option{}
	list.Position = 0
	list.Length = 0
}

func set(options []model.Option, title string) {
	list.Options = options
	list.Title = title
	list.Position = 0
	list.Length = len(options)
}
