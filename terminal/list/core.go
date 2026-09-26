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

func set(payload model.ListPayload) {
	list.Options = payload.Options
	list.Title = payload.Title
	list.Position = 0
	list.Length = len(payload.Options)
}
