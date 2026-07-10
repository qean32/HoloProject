package menu

import (
	"main/deep"
	"main/lib"
	"main/lib/array"
	"main/model"
	"main/terminal/field"
	"main/terminal/list"
)

func runCommand() {
	list.List(
		array.Map(deep.TMP_COMMANDS, func(value []string) model.Option {
			return model.Option{
				Message: value[0],
				Command: func() {
					lib.RunCommand(value[1])
				},
			}
		}),
	)
}

func manual() {
	field.Field()
}

func addCommand() {
}
