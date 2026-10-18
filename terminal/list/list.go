package list

import (
	"fmt"
	"strconv"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"main/callstack"
	"main/callstack/event"
	"main/constants"
	"main/constants/sgr"
	"main/lib/utils"
	"main/model"
	"main/terminal"
)

type state struct {
	position int
	options  []model.Option
}

func (s *state) length() int {
	return len(s.options)
}

func List(payload model.ListPayload) {
	terminal.DownAndStart()
	if len(payload.Options) == 0 {
		event.Response("length == zero! ave caesar! ")
		return
	}

	cursor.Hide()
	terminal.Println(utils.GetCustomMessage(payload.Title+" ↓"+" (ctrl+c back)", constants.StyleDim...))
	terminal.PrintTopLine()

	s := &state{options: payload.Options}
	s.render()
	s.jumpToStart()

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if value, err := strconv.Atoi(key.String()); err == nil {
			s.jumpTo(value - 1)
		}
		switch key.Code {
		case keys.Enter:
			s.selectOption()
			return true, nil
		case keys.Down:
			s.moveDown()
		case keys.Up:
			s.moveUp()
		case keys.Escape, keys.CtrlC:
			cursor.Show()
			s.jumpToEnd()
			return true, nil
		}

		return false, nil
	})
	terminal.DownAndStart()
}

func (s *state) render() {
	for i, item := range s.options {
		isSelected := i == s.position

		color := sgr.DIM
		if isSelected {
			color = sgr.GREEN
		}
		startChar := utils.GetCustomMessage(getStartSymbol(isSelected), color)
		label := utils.GetCustomMessage(fmt.Sprintf("%d. %s", i+1, item.Message), sgr.DIM)

		cursor.ClearLine()
		if isSelected {
			terminal.Print(utils.GetCustomMessage("│", sgr.DIM), startChar+label)
		} else {
			terminal.Print(startChar + label)
		}
		terminal.DownAndStart()
	}
}

func (s *state) rerender() {
	s.jumpToEnd()
	terminal.ClearLines(s.length())
	s.render()
	s.jumpToStart()
}

func (s *state) jumpTo(newPos int) {
	if newPos < 0 || newPos >= s.length() {
		return
	}
	diff := newPos - s.position
	if diff > 0 {
		cursor.Down(diff)
	} else if diff < 0 {
		cursor.Up(-diff)
	}
	s.position = newPos
	s.rerender()
}

func (s *state) moveUp() {
	if s.decrementPosition() {
		s.rerender()
	}
}

func (s *state) moveDown() {
	if s.incrementPosition() {
		s.rerender()
	}
}

func (s *state) incrementPosition() bool {
	if s.position < s.length()-1 {
		s.position++
		cursor.Down(1)
		return true
	}
	if s.length() > 1 {
		s.position = 0
		s.jumpToStart()
		cursor.Down(1)
		return true
	}
	return false
}

func (s *state) decrementPosition() bool {
	if s.position > 0 {
		s.position--
		cursor.Up(1)
		return true
	}
	if s.length() > 1 {
		diff := s.length() - 1 - s.position
		if diff > 0 {
			cursor.Down(diff)
		}
		s.position = s.length() - 1
		return true
	}
	return false
}

func (s *state) selectOption() {
	s.jumpToEnd()
	cursor.Show()
	callstack.Push(utils.GetShortEvent(s.options[s.position].Event))
}

func (s *state) jumpToEnd() {
	cursor.Down(s.length() - s.position)
}

func (s *state) jumpToStart() {
	cursor.Up(s.length() - s.position)
}

func getStartSymbol(isSelected bool) string {
	if isSelected {
		return constants.CHAR_SELECTED_ITEM
	}
	return constants.CHAR_UN_SELECTED_ITEM
}
