package main

import (
	"main/callstack"
	"main/lib"
	"main/lib/manual"
	"main/terminal"
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
			terminal.OutputTechInfo(event)
			lib.Event(event)
		}
		manual.Manual()
		<-callstack.Changed()
	}
}
