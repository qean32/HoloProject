package list

import (
	"fmt"
	"main/lib"
	"main/model"
	"main/terminal"
	"strconv"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func List(_list []model.Option) {
	cursor.Hide()
	set(_list)
	renderList(_list)
	cursor.Up(4)
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
		fmt.Println(getStartChar(i == list.Position) + item.Message)
		cursor.StartOfLineDown(0)
	}
}

func selectItem() {
	jumpToEndList()
	list.Options[list.Position].Command()
	cursor.Show()
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
	cursor.Up(list.Length - list.Position)
}
