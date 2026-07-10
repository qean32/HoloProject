package list

import (
	"main/constants"
	"main/constants/literals"
	"main/lib"
	"main/model"
	"main/terminal"
	"strconv"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func List(options []model.Option) {
	cursor.Hide()
	set(options)
	renderList(options)
	jumpToStartList()
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if value, err := strconv.Atoi(key.String()); err == nil {
			list.Position = value
			reRenderList()
		}
		switch key.Code {
		case keys.Enter:
			selectItem()
			reset()
			return true, nil
		case keys.Down:
			moveDown()
		case keys.Up:
			moveUp()
		case keys.Escape:
			lib.StopProcess()
		case keys.CtrlC:
			lib.StopProcess()
		}
		return false, nil
	})
}

func renderList(options []model.Option) {
	for i, item := range options {
		isSelected := i == list.Position

		if isSelected {
			terminal.Output(terminal.GetCustomMessage(getStartChar(isSelected), literals.SGR.GREEN) +
				terminal.GetCustomMessage(item.Message, constants.StyleError...))

		} else {
			terminal.Output(terminal.GetCustomMessage(getStartChar(isSelected), literals.SGR.GREEN) +
				terminal.GetCustomMessage(item.Message, literals.SGR.DIM, literals.SGR.WHITE))

		}
		terminal.DownAndStart()
	}
}

func selectItem() {
	jumpToEndList()
	cursor.Show()
	list.Options[list.Position].Command()
}

func moveUp() {
	if decrimentPosition() {
		reRenderList()
	}
}

func moveDown() {
	if incrementPosition() {
		reRenderList()
	}
}

func reRenderList() {
	jumpToEndList()
	terminal.ClearLines(list.Length)
	renderList(list.Options)
	jumpToStartList()
}
