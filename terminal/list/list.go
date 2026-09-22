package list

import (
	"fmt"
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
	terminal.OutputTechInfo("[LIST] start")
	defer reset()

	cursor.Hide()
	set(options)
	renderList(options)
	jumpToStartList()

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if value, err := strconv.Atoi(key.String()); err == nil {
			newPos := value - 1
			if newPos < 0 || newPos >= list.Length {
				return false, nil
			}
			diff := newPos - list.Position
			if diff > 0 {
				cursor.Down(diff)
			} else if diff < 0 {
				cursor.Up(-diff)
			}
			list.Position = newPos
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
			return true, nil
		case keys.CtrlC:
			low.StopProcess()
			return true, nil
		}

		return false, nil
	})
	terminal.OutputTechInfo("[LIST] listener returned")
}

func renderList(options []model.Option) {
	for i, item := range options {
		isSelected := i == list.Position
		startChar := getStartChar(isSelected)

		style := literals.SGR.DIM
		if isSelected {
			startChar = terminal.GetCustomMessage(startChar, literals.SGR.GREEN)
		} else {
			startChar = terminal.GetCustomMessage(startChar, literals.SGR.DIM)
		}

		cursor.ClearLine()
		terminal.Output(startChar + terminal.GetCustomMessage(fmt.Sprintf("%d. %s", i+1, item.Message), style))
		terminal.DownAndStart()
	}
}

func _select(event model.Event) {
	jumpToEndList()
	cursor.Show()
	terminal.DownAndStart()
	callstack.PushCallStack(event)
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
