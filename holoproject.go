package main

import (
	"fmt"
	"runtime"

	"main/callstack"
	"main/callstack/event"
	eventlist "main/constants/event-list"
	"main/lib"
	"main/terminal"
	"main/terminal/qcallback"
)

var logger = true

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
			if logger && ev.Key != eventlist.RESPONSE {
				terminal.PrintTechInfo(ev)
			}
			if err := lib.HandleEvent(ev); err != nil {
				callstack.PushError(ev, err)
			}
		}
		qcallback.QCallback(func() { event.Reset() })
		<-callstack.Changed()
	}
}
