package main

import (
	"main/callstack"
	"main/constants/literals"
	"main/lib"
	"main/model"
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
		lib.Event(model.Event{Key: literals.COMMANDS_LIST.MANUAL})
		<-callstack.Changed()
	}
}
