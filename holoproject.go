package main

import (
	"main/lib"
	"main/model"
)

func main() {
	callstack := make(model.Channel)
	go lib.INIT(callstack)
	LOOP(callstack)
}

func LOOP(callstack model.Channel) {
	for value := range callstack {
		lib.ITERATION_CYCLE(value)
		// fmt.Print(value)
	}
}
