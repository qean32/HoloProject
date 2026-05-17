package main

import (
	"main/deep"
	"main/lib"
)

func main() {
	go lib.INIT()
	LOOP()
}

func LOOP() {
	for value := range deep.Callstack_channel {
		lib.ITERATION_CYCLE(value)
	}
}
