package list

import (
	"fmt"
	"main/lib"
	"main/model"
	"main/terminal"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func List(list []model.Option) {
	set(list)
	renderList(list)
	terminal.VerticalJumpUpCount(len(list))
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.Enter:
			selectItem()
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
		fmt.Println("\r" + getStartChar(i == list.Position) + item.Message + "\r")
	}
}

func selectItem() {
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
	terminal.VerticalJumpDownCount(list.Length - list.Position)
	terminal.ClearLines(list.Length)
	renderList(list.Options)
	terminal.VerticalJumpUpCount(list.Length - list.Position)
}
