package qnext

import (
	"main/callstack"
	eventlist "main/constants/event-list"
	"main/lib/death"
	"main/model"
	"main/terminal"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func QNext() {
	terminal.Print("Press Enter to continue")
	cursor.Hide()
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.Enter:
			cursor.Show()
			callstack.PushCallStack(model.Event{Key: eventlist.RESET})
			return true, nil
		case keys.Escape, keys.CtrlC:
			cursor.Show()
			death.Exit()
			return true, nil
		}

		return false, nil
	})
}
