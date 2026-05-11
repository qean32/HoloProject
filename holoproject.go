package main

import (
	"main/lib"
	"main/model"
)

func main() {
	callstack_channel := make(model.CallStackChannel)
	lib.INIT(callstack_channel)
	// LOOP(callstack_channel)
}

func LOOP(callstack_channel model.CallStackChannel) {
	for value := range callstack_channel {
		lib.ITERATION_CYCLE(value)
	}
}
