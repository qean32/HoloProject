package main

import (
	"main/callstack"
	"main/constants"
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
			if constants.Mode != "prod" {
				terminal.PrintTechInfo(event)
			}
			lib.Event(event)
		}
		qnext.QNext()
		<-callstack.Changed()
	}
}
