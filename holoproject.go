package main

import (
	"fmt"
	"runtime"

	"main/callstack"
	eventlist "main/constants/event-list"
	"main/lib"
	"main/model"
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

var logger = false

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
			lib.HandleEvent(event)
		}
		qcallback.QCallback(func() { callstack.Push(model.Event{Key: eventlist.RESET}) })
		<-callstack.Changed()
	}
}
