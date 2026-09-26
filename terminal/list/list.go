package list

import (
	"fmt"
	"strconv"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"main/callstack"
	"main/constants/literals"
	"main/lib/death"
	"main/lib/utils"
	"main/model"
	"main/terminal"
)

func List(payload model.ListPayload) {
	defer reset()
	if len(payload.Options) == 0 {
		return
	}
	cursor.Hide()
	terminal.Println(payload.Title)
	terminal.PrintTopLine()

	set(payload)
	renderList(payload.Options)
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
			_select()
			return true, nil
		case keys.Down:
			moveDown()
		case keys.Up:
			moveUp()
		case keys.Escape, keys.CtrlC:
			jumpToEndList()
			return true, nil
		}

		return false, nil
	})
}

func renderList(options []model.Option) {
	for i, item := range options {
		isSelected := i == list.Position
		startChar := getStartSymbol(isSelected)

		style := literals.SGR.DIM
		color := literals.SGR.DIM
		if isSelected {
			color = literals.SGR.GREEN
		}
		startChar = utils.GetCustomMessage(startChar, color)

		cursor.ClearLine()
		if isSelected {
			terminal.Print(utils.GetCustomMessage("│", literals.SGR.DIM), startChar+utils.GetCustomMessage(fmt.Sprintf("%d. %s", i+1, item.Message), style))
		} else {
			terminal.Print(startChar + utils.GetCustomMessage(fmt.Sprintf("%d. %s", i+1, item.Message), style))
		}
		terminal.DownAndStart()
	}
}

func _select() {
	jumpToEndList()
	terminal.DownAndStart()
	cursor.Show()
	callstack.PushCallStack(death.GetShortEvent(list.Options[list.Position].Event))
}

func reRenderList() {
	jumpToEndList()
	terminal.ClearLines(list.Length)
	renderList(list.Options)
	jumpToStartList()
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

func moveUp() {
	if decrementPosition() {
		reRenderList()
	}
}

func moveDown() {
	if incrementPosition() {
		reRenderList()
	}
}
