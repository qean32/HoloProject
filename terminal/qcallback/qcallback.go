package qcallback

import (
	"main/callstack/event"
	"main/terminal"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func QCallback(callback func()) {
	terminal.Print("press Enter to continue")
	cursor.Hide()
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.Enter:
			cursor.Show()
			callback()
			return true, nil
		case keys.Escape, keys.CtrlC:
			event.Reset()
			return true, nil
		}

		return false, nil
	})
}
