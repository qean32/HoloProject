package main

import (
	"main/callstack"
	"main/callstack/manual"
	"main/lib"
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
			terminal.PrintTechInfo(event)
			lib.Event(event)
		}
		manual.Manual()
		<-callstack.Changed()
	}
}
