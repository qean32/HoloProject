package main

import (
	"main/callstack"
	"main/lib"
	"main/terminal"
)

func main() {
	go RunLoop()
	lib.INIT()

}

func RunLoop() {
	for {
		for {
			event, ok := callstack.GetTask()
			if !ok {
				break
			}
			terminal.OutputTechInfo(event.Key)
			lib.Event(event)
		}
		<-callstack.Changed()
	}
}
