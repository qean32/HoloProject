package main

import (
	test "main/__test__"
	"main/callstack"
	"main/lib"
	"main/terminal"
	qnext "main/terminal/qNext"
)

func main() {
	lib.Init()
	test.Test()
	go RunLoop()

	select {}
}

var log = true

func RunLoop() {
	for {
		for {
			event, ok := callstack.Pop()
			if !ok {
				break
			}
			if log {
				terminal.PrintTechInfo(event)
			}
			lib.HandleEvent(event)
		}
		qnext.QNext()
		<-callstack.Changed()
	}
}
