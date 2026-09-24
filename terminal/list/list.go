package list

import (
	"fmt"
	"strconv"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"main/callstack"
	"main/constants/literals"
	"main/lib/low"
	"main/model"
	"main/terminal"
)

func List(options []model.Option, title string) {
	if len(options) == 0 {
		return
	}
	terminal.Println(title)
	terminal.TopLine()
	defer reset()

	cursor.Hide()
	set(options, title)
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
			selectOption(options[list.Position].Event)
			return true, nil
		case keys.Down:
			moveDown()
		case keys.Up:
			moveUp()
		case keys.Escape, keys.CtrlC:
			return true, nil
		}

		return false, nil
	})
}

func renderList(options []model.Option) {
	for i, item := range options {
		isSelected := i == list.Position
		startChar := getStartChar(isSelected)

		style := literals.SGR.DIM
		color := literals.SGR.DIM
		if isSelected {
			color = literals.SGR.GREEN
		}
		startChar = terminal.GetCustomMessage(startChar, color)

		cursor.ClearLine()
		terminal.Print(startChar + terminal.GetCustomMessage(fmt.Sprintf("%d. %s", i+1, item.Message), style))
		terminal.DownAndStart()
	}
}

func selectOption(event model.Event) {
	jumpToEndList()
	terminal.DownAndStart()
	cursor.Show()
	callstack.PushCallStack(low.GetShortEvent(event))
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

func reRenderList() {
	jumpToEndList()
	terminal.ClearLines(list.Length)
	renderList(list.Options)
	jumpToStartList()
}
