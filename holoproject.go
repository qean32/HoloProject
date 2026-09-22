package main

import (
	"main/callstack"
	"main/lib"
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
			lib.Event(event)
		}
		<-callstack.Changed()
	}
}
