package qnext

import (
	"main/callstack"
	"main/constants/literals"
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
			callstack.PushCallStack(model.Event{Key: literals.EventList.RESET})
			return true, nil
		case keys.Escape, keys.CtrlC:
			cursor.Show()
			death.Exit()
			return true, nil
		}

		return false, nil
	})
}
