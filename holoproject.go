package main

import (
	"main/callstack"
	"main/lib"
	"main/terminal"
	qnext "main/terminal/qNext"
)

func main() {
	lib.INIT()
	go RunLoop()

	select {}
}

func RunLoop() {
	for {
		for {
			event, ok := callstack.Pop()
			if !ok {
				break
			}
			if true {
				terminal.PrintTechInfo(event)
			}
			lib.Event(event)
		}
		qnext.QNext()
		<-callstack.Changed()
	}
}
