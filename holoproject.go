package main

import (
	"main/callstack"
	"main/lib"
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

			lib.Event(event)
		}

		<-callstack.Changed()
	}
}
