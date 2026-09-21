package list

import (
	"main/callstack"
	"main/constants/literals"
	"main/lib/low"
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
			_select(options[list.Position].Event)
			return true, nil
		case keys.Down:
			moveDown()
		case keys.Up:
			moveUp()
		case keys.Escape:
			low.StopProcess()
		case keys.CtrlC:
			low.StopProcess()
		}

		return false, nil
	})
}

func renderList(options []model.Option) {
	for i, item := range options {
		isSelected := i == list.Position

		if isSelected {
			terminal.Output(terminal.GetCustomMessage(getStartChar(isSelected), literals.SGR.GREEN) +
				terminal.GetCustomMessage(item.Message)) // constants.StyleSelected...,
		} else {
			terminal.Output(terminal.GetCustomMessage(getStartChar(isSelected), literals.SGR.GREEN) +
				terminal.GetCustomMessage(item.Message, literals.SGR.DIM))

		}
		terminal.DownAndStart()
	}
}

func _select(event model.Event) {
	jumpToEndList()
	cursor.Show()
	terminal.DownAndStart()
	callstack.PushCallStack(event)
	reset()
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
