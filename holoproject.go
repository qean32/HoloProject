package main

import (
	"fmt"
	"runtime"

	"main/callstack"
	"main/callstack/event"
	eventlist "main/constants/event-list"
	"main/lib"
	"main/terminal"
	qcallback "main/terminal/qcallback"
)

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("Эта программа поддерживается только на Windows.")
		return
	}

	lib.Init()
	go RunLoop()

	select {}
}

var logger = true

func RunLoop() {
	for {
		for {
			event, ok := callstack.Pop()
			if !ok {
				break
			}
			if logger && event.Key != eventlist.RESPONSE {
				terminal.PrintTechInfo(event)
			}
			if err := lib.HandleEvent(event); err != nil {
				callstack.PushError(event, err)
			}
		}
		qcallback.QCallback(func() { event.Reset() })
		<-callstack.Changed()
	}
}
