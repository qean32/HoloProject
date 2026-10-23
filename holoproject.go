package main

import (
	"fmt"
	"runtime"

	"main/callstack"
	eventpkg "main/callstack/event"
	eventlist "main/constants/event-list"
	"main/lib"
	"main/lib/eventlog"
	"main/terminal"
	"main/terminal/qcallback"
)

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("Эта программа поддерживается только на Windows.")
		return
	}

	lib.Init()
	go runLoop()

	select {}
}

func runLoop() {
	for {
		for {
			ev, ok := callstack.Pop()
			if !ok {
				break
			}
			if eventlog.Enabled() && ev.Key != eventlist.RESPONSE {
				terminal.PrintTechInfo(ev)
			}
			if err := lib.HandleEvent(ev); err != nil {
				callstack.PushError(ev, err)
			}
		}
		qcallback.QCallback(func() { eventpkg.Reset() })
	}
}
